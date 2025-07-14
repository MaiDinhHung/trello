<template>
  <div class="min-h-screen bg-gray-100">
    <!-- Navbar -->
    <nav class="flex justify-between items-center px-6 py-4 bg-white shadow-sm">
      <h1 class="text-2xl font-bold text-primary">
        {{ board?.title || "Board" }}
      </h1>
      <div class="flex items-center gap-3">
        <span>{{ user?.name }}</span>
        <UButton color="red" @click="logout">Đăng xuất</UButton>
      </div>
    </nav>

    <!-- Board content -->
    <div class="p-4">
      <div class="flex gap-4 overflow-x-auto">
        <UCard
          v-for="list in lists"
          :key="list.id"
          class="w-64 min-h-[300px] flex-shrink-0 bg-white shadow p-4 flex flex-col justify-between"
        >
          <div class="flex justify-between items-center mb-2">
            <div v-if="!list.editing" class="flex justify-between w-full">
              <h3 class="font-extrabold text-xl">{{ list.title }}</h3>
              <div class="flex gap-1">
                <button
                  class="text-xs text-blue-500"
                  @click="list.editing = true"
                >
                  Sửa
                </button>
                <button
                  class="text-xs text-red-500"
                  @click="deleteList(list.id)"
                >
                  Xóa
                </button>
              </div>
            </div>
            <div v-else class="flex gap-2 w-full">
              <UInput v-model="list.title" size="xs" class="w-full" />
              <UButton size="xs" @click="updateList(list)">Lưu</UButton>
              <UButton size="xs" color="gray" @click="list.editing = false"
                >Hủy</UButton
              >
            </div>
          </div>

          <!-- Thẻ (cards) -->
          <draggable
            v-model="list.cards"
            group="cards"
            item-key="id"
            :data-list-id="list.id"
            @end="(event) => onCardDrop(event, list.id)"
            class="flex flex-col gap-2 mb-2"
          >
            <template #item="{ element }">
              <div
                class="bg-gray-100 p-2 rounded shadow cursor-pointer flex justify-start items-center"
                @click="openCard(element)"
              >
              <!-- <UCheckbox 
              v-model="element.completed"
              @change = "toogleCheckcard(element)"
                  class="float-right pe-1.5"
                  size="xs"
                
              /> -->
                {{ element.title }}
              </div>
            </template>
          </draggable>

          <!-- Thêm thẻ -->
          <div v-if="list.showAddCard">
            <UInput
              v-model="list.newCardTitle"
              placeholder="Tên thẻ..."
              class="mb-2"
            />
            <div class="flex gap-2">
              <UButton size="xs" @click="createCard(list)">Thêm</UButton>
              <UButton size="xs" color="gray" @click="list.showAddCard = false"
                >Hủy</UButton
              >
            </div>
          </div>
          <div v-else>
            <button
              class="text-sm text-gray-500 hover:text-primary"
              @click="list.showAddCard = true"
            >
              + Thêm thẻ
            </button>
          </div>
        </UCard>

        <!-- Thêm cột mới -->
        <UCard
          v-if="!showCreateInput"
          class="w-64 h-96 flex items-center justify-center flex-shrink-0 border-dashed border-2 border-gray-300 hover:border-primary hover:text-primary cursor-pointer"
          @click="showCreateInput = true"
        >
          <span>+ Tạo cột mới</span>
        </UCard>

        <UCard
          v-else
          class="w-64 h-96 flex flex-col p-4 justify-between flex-shrink-0 border border-gray-300"
        >
          <div>
            <UInput
              v-model="newListTitle"
              placeholder="Tên cột..."
              class="mb-2"
            />
          </div>
          <div class="flex gap-2">
            <UButton @click="createList" color="primary">Thêm</UButton>
            <UButton @click="showCreateInput = false" color="gray">Hủy</UButton>
          </div>
        </UCard>
      </div>
    </div>
  </div>
  <UModal
    v-model:open="showCardModal"
    :title="selectedCard?.title"
    :description="`Chi tiết thẻ công việc`"
    :close="{ color: 'gray', class: 'rounded-full' }"
    class="max-w-4xl"
  >
    <template #body>
      <div class="flex flex-col md:flex-row gap-6">
        <!-- Bên trái - Thông tin thẻ -->
        <div class="flex-1 space-y-4">
          <!-- Thành viên, nhãn, ngày -->
          <div class="flex flex-wrap items-center gap-3">
            <div class="flex items-center gap-2">
              <span class="font-semibold">Thành viên:</span>
              <div
                class="rounded-full bg-blue-600 text-white w-8 h-8 flex items-center justify-center font-bold"
              >
                MH
              </div>
              <UButton
                size="xs"
                icon="i-lucide-plus"
                color="gray"
                variant="soft"
              />
            </div>

            <div class="flex items-center gap-2">
              <span class="font-semibold">Nhãn:</span>
              <span class="w-6 h-6 bg-green-500 rounded" title="a"></span>
              <span class="w-6 h-6 bg-yellow-400 rounded" title="b"></span>
              <UButton
                size="xs"
                icon="i-lucide-plus"
                color="gray"
                variant="soft"
              />
            </div>

            <!-- Ngày -->
            <UButton
              icon="i-lucide-calendar"
              @click="showDateModal = true"
              class="flex items-center gap-2 border-1"
              color="none"
            >
              {{ formattedDates }}

              <template v-if="dateStatus === 'soon'">
                <span
                  class="px-2 py-0.5 bg-yellow-400 text-white text-xs rounded"
                  >Sắp tới hạn</span
                >
              </template>
              <template v-else-if="dateStatus === 'overdue'">
                <span class="px-2 py-0.5 bg-red-500 text-white text-xs rounded"
                  >Quá hạn</span
                >
              </template>
            </UButton>

            <UModal
              v-model:open="showDateModal"
              title="Ngày"
              class="max-w-lg h-[600px]"
            >
              <template #body>
                <div class="">
                  <div class="space-y-2">
                    <label class="font-semibold">Ngày bắt đầu</label>
                    <Datepicker
                      v-model="tempStartDate"
                      placeholder="Chọn ngày bắt đầu"
                    />
                  </div>
                  <div class="space-y-2">
                    <label class="font-semibold">Ngày hết hạn</label>
                    <Datepicker
                      v-model="tempEndDate"
                      placeholder="Chọn ngày hết hạn"
                    />
                  </div>

                  <div class="flex gap-2 pt-5">
                    <UButton @click="confirmDates" color="primary">Lưu</UButton>
                    <UButton @click="removeDates" color="gray">Gỡ bỏ</UButton>
                  </div>
                </div>
              </template>
            </UModal>
          </div>

          <!-- Mô tả -->
          <div>
            <h3 class="font-semibold flex items-center gap-2 mt-4">
              <i class="i-lucide-align-left" /> Mô tả
            </h3>

            <div
              v-if="!editingDescription"
              class="flex justify-between items-start gap-4"
            >
              <p class="text-gray-700 whitespace-pre-wrap flex-1 ps-2.5">
                {{ selectedCard?.description || "Không có mô tả" }}
              </p>
              <UButton
                size="xs"
                color="gray"
                variant="soft"
                @click="startEditDescription"
                >Chỉnh sửa</UButton
              >
            </div>

            <div v-else class="flex flex-col gap-2">
              <UTextarea
                v-model="descriptionInput"
                placeholder="Nhập mô tả..."
                :rows="2"
              />
              <div class="flex gap-2">
                <UButton size="xs" color="primary" @click="saveDescription"
                  >Lưu</UButton
                >
                <UButton size="xs" color="gray" @click="cancelEditDescription"
                  >Hủy</UButton
                >
              </div>
            </div>
          </div>

          <!-- Checklist -->
          <div>
            <h3 class="font-semibold flex items-center gap-2 mt-4">
              <i class="i-lucide-check-square" /> Việc cần làm
            </h3>

            <div
              v-for="item in checklist"
              :key="item.id"
              class="flex items-center gap-2 mt-2"
            >
              <UCheckbox
                v-model="item.completed"
                @change="toggleChecklist(item)"
                class="ps-2.5"
              />
              <div v-if="item.editing">
                <UInput v-model="item.content" size="xs" />
                <UButton size="xs" @click="saveChecklist(item)">Lưu</UButton>
                <UButton size="xs" color="gray" @click="item.editing = false"
                  >Hủy</UButton
                >
              </div>

              <div v-else class="flex items-center gap-2">
                <span
                  :class="{ 'line-through text-gray-500': item.completed }"
                  >{{ item.content }}</span
                >
                <UButton
                  size="xs"
                  icon="i-lucide-pencil"
                  @click="item.editing = true"
                  color="none"
                />
              </div>

              <UButton
                icon="i-lucide-trash"
                color="error"
                variant="ghost"
                size="xs"
                @click="deleteChecklist(item.id)"
              />
            </div>

            <div class="flex gap-2 mt-2">
              <UInput
                v-model="newChecklist"
                placeholder="Thêm mục mới..."
                class="w-fit"
              />
              <UButton size="xs" @click="addChecklist">Thêm</UButton>
            </div>
          </div>
        </div>
        <!-- Ngăn cách -->
        <div class="hidden md:block w-px bg-gray-300"></div>
        <!-- Bên phải: Bình luận -->
        <div class="md:w-1/4 w-full pl-4">
          <h3 class="font-semibold">Nhận xét</h3>
          <div
            v-for="comment in comments"
            :key="comment.id"
            class="text-sm mb-2"
          >
            <span class="text-gray-500">{{ comment.created_at }}</span
            ><br />
            {{ comment.content }}
          </div>
          <div class="mt-2 flex gap-2">
            <UInput v-model="newComment" placeholder="Viết bình luận..." />
            <UButton @click="addComment" size="xs">Gửi</UButton>
          </div>
        </div>
      </div>
    </template>
  </UModal>
</template>

<script setup>
import { ref, onMounted, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import draggable from "vuedraggable";

// --- State ---
const route = useRoute();
const router = useRouter();
const boardId = route.params.id;

const user = ref(null);
const board = ref({});
const lists = ref([]);

const newListTitle = ref("");
const showCreateInput = ref(false);

const showCardModal = ref(false);
const selectedCard = ref(null);

const comments = ref([]);
const newComment = ref("");

const checklist = ref([]);
const newChecklist = ref("");

const editingDescription = ref(false);
const descriptionInput = ref("");

// --- Auth ---
function logout() {
  if (!confirm("Bạn có chắc muốn đăng xuất không?")) return;
  localStorage.clear();
  router.push("/login");
}

// --- Fetch board + list ---
onMounted(async () => {
  const saved = localStorage.getItem("user");
  if (!saved) return router.push("/login");
  user.value = JSON.parse(saved);

  const { data: boardData } = await useFetch(
    `http://localhost:3001/api/boards/${boardId}`
  );
  const { data: listData } = await useFetch(
    `http://localhost:3001/api/boards/${boardId}/lists`
  );

  board.value = boardData.value;
  lists.value = (listData.value || []).map((list) => ({
    ...list,
    showAddCard: false,
    newCardTitle: "",
    editing: false,
    cards: (list.cards || []).sort((a, b) => a.position - b.position),
  }));
});

// --- List ---
async function createList() {
  if (!newListTitle.value.trim()) return;

  const { data, error } = await useFetch("http://localhost:3001/api/lists", {
    method: "POST",
    body: {
      title: newListTitle.value,
      board_id: parseInt(boardId),
    },
  });

  if (error.value) return alert("Tạo cột thất bại");

  lists.value.push({
    ...data.value,
    showAddCard: false,
    newCardTitle: "",
    editing: false,
    cards: [],
  });

  newListTitle.value = "";
  showCreateInput.value = false;
}

async function updateList(list) {
  try {
    await $fetch(`http://localhost:3001/api/lists/${list.id}`, {
      method: "PUT",
      body: { title: list.title },
    });
    list.editing = false;
  } catch {
    alert("Không thể cập nhật cột");
  }
}

async function deleteList(id) {
  if (!confirm("Bạn có chắc muốn xoá cột này không?")) return;
  try {
    await $fetch(`http://localhost:3001/api/lists/${id}`, {
      method: "DELETE",
    });
    lists.value = lists.value.filter((l) => l.id !== id);
  } catch {
    alert("Không thể xoá cột");
  }
}

// --- Card ---
async function createCard(list) {
  const position =
    list.cards.length > 0 ? list.cards[list.cards.length - 1].position + 1 : 0;
  const payload = {
    title: list.newCardTitle,
    list_id: list.id,
    position,
    description: "",
    start_date: "",
    end_date: "",
    user_id: user.value?.id || 1, //// cần fix lại
  };

  const { data, error } = await useFetch("http://localhost:3001/api/cards", {
    method: "POST",
    body: payload,
  });

  if (error.value) return alert("Không thể tạo thẻ.");

  list.cards.push(data.value);
  list.newCardTitle = "";
  list.showAddCard = false;
}

// async function toggleCheckcard(item) {
//   item.completed = !item.completed;
//   await $fetch(`http://localhost:3001/api/cards/${item.id}`, {
//     method: "PUT",
//     body: { completed: item.completed },
//   });
// }

async function onCardDrop(event, toListId) {
  const toList = lists.value.find((list) => list.id === toListId);
  if (!toList) return;

  await Promise.all(
    toList.cards.map((card, i) =>
      $fetch(`http://localhost:3001/api/cards/${card.id}`, {
        method: "PUT",
        body: { list_id: toListId, position: i },
      })
    )
  );
}

// --- Modal open card---
function openCard(card) {
  selectedCard.value = card;
  showCardModal.value = true;
  fetchComments(card.id);
  fetchChecklist();
}

watch(showCardModal, (open) => {
  if (!open) {
    checklist.value = [];
    comments.value = [];
    selectedCard.value = null;
  }
});

// --- Comment ---
async function fetchComments(cardId) {
  const { data } = await useFetch(
    `http://localhost:3001/api/cards/${cardId}/comments`
  );
  comments.value = data.value || [];
}

async function addComment() {
  if (!newComment.value.trim()) return;
  await $fetch(
    `http://localhost:3001/api/cards/${selectedCard.value.id}/comments`,
    {
      method: "POST",
      body: {
        content: newComment.value,
        user_id: user.value?.id || 1,
      },
    }
  );
  newComment.value = "";
  fetchComments(selectedCard.value.id);
}

// --- Checklist ---
async function fetchChecklist() {
  const { data } = await useFetch(
    `http://localhost:3001/api/cards/${selectedCard.value.id}/checklist`
  );
  checklist.value = data.value || [];
}

async function addChecklist() {
  if (!newChecklist.value.trim()) return;
  const res = await $fetch(
    `http://localhost:3001/api/cards/${selectedCard.value.id}/checklist`,
    {
      method: "POST",
      body: {
        content: newChecklist.value,
        completed: false,
      },
    }
  );
  newChecklist.value = "";
  await fetchChecklist();
}
async function saveChecklist(item) {
  await $fetch(`http://localhost:3001/api/checklist/${item.id}`, {
    method: "PUT",
    body: {
      content: item.content,
      completed: item.completed,
    },
  });
  item.editing = false;
  await fetchChecklist();
}

async function toggleChecklist(item) {
  await $fetch(`http://localhost:3001/api/checklist/${item.id}`, {
    method: "PUT",
    body: {
      content: item.content,
      completed: item.completed,
    },
  });
  await fetchChecklist();
}

async function deleteChecklist(id) {
  if (!confirm("Bạn có chắc muốn xoá mục này không?")) return;
  await $fetch(`http://localhost:3001/api/checklist/${id}`, {
    method: "DELETE",
  });
  await fetchChecklist();
}

// description

function startEditDescription() {
  descriptionInput.value = selectedCard.value.description || "";
  editingDescription.value = true;
}

function cancelEditDescription() {
  editingDescription.value = false;
}

async function saveDescription() {
  await $fetch(`http://localhost:3001/api/cards/${selectedCard.value.id}`, {
    method: "PUT",
    body: {
      description: descriptionInput.value,
    },
  });
  selectedCard.value.description = descriptionInput.value;
  editingDescription.value = false;
}

// Date
import Datepicker from "@vuepic/vue-datepicker";
import "@vuepic/vue-datepicker/dist/main.css";
const startDate = ref(null);
const endDate = ref(null);
const showDateModal = ref(false);
const tempStartDate = ref(null);
const tempEndDate = ref(null);

watch(showCardModal, (open) => {
  if (open) {
    tempStartDate.value = selectedCard.value.start_date
      ? new Date(selectedCard.value.start_date)
      : null;
    tempEndDate.value = selectedCard.value.end_date
      ? new Date(selectedCard.value.end_date)
      : null;
  } else {
    tempStartDate.value = null;
    tempEndDate.value = null;
  }
});
async function confirmDates() {
  await $fetch(`http://localhost:3001/api/cards/${selectedCard.value.id}`, {
    method: "PUT",
    body: {
      start_date: tempStartDate.value
        ? tempStartDate.value.toISOString()
        : null,
      end_date: tempEndDate.value ? tempEndDate.value.toISOString() : null,
    },
  });

  selectedCard.value.start_date = tempStartDate.value;
  selectedCard.value.end_date = tempEndDate.value;
  showDateModal.value = false;
}

async function removeDates() {
  await $fetch(`http://localhost:3001/api/cards/${selectedCard.value.id}`, {
    method: "PUT",
    body: {
      start_date: null,
      end_date: null,
    },
  });

  selectedCard.value.start_date = null;
  selectedCard.value.end_date = null;
  tempStartDate.value = null;
  tempEndDate.value = null;
  showDateModal.value = false;
}

const formattedDates = computed(() => {
  if (!selectedCard.value?.start_date || !selectedCard.value?.end_date)
    return "";

  const start = new Date(selectedCard.value.start_date);
  const end = new Date(selectedCard.value.end_date);

  const startDateStr = start.toLocaleDateString("vi-VN", {
    day: "2-digit",
    month: "2-digit",
  });
  const startTimeStr = start.toLocaleTimeString("vi-VN", {
    hour: "2-digit",
    minute: "2-digit",
  });

  const endDateStr = end.toLocaleDateString("vi-VN", {
    day: "2-digit",
    month: "2-digit",
  });

  return `${startDateStr} || ${startTimeStr}   ${endDateStr}`;
});

const dateStatus = computed(() => {
  if (!selectedCard.value?.end_date) return "normal";

  const now = new Date();
  const end = new Date(selectedCard.value.end_date);

  if (now > end) return "overdue";

  const diff = (end - now) / (1000 * 60 * 60 * 24);
  if (diff <= 1) return "soon";

  return "normal";
});
</script>
