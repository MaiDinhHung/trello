<template>
  <div class="min-h-screen bg-white">
    <!-- Navbar -->
    <nav
      class="flex justify-between items-center px-6 py-4 bg-blue-300 shadow-sm h-23"
    >
      <h1 class="text-2xl font-bold text-black">
        {{ board?.title || "Board" }}
      </h1>
      <div class="flex items-center gap-3">
        <UPopover
          v-model:open="open"
          :dismissible="false"
          :ui="{ content: 'p-4' }"
        >
          <UButton icon="i-lucide-filter" color="neutral" variant="subtle" />

          <template #content>
            <div class="flex items-center justify-between gap-4 mb-6">
              <h2 class="text-highlighted font-semibold">
                Bộ lọc
              </h2>
              <UButton
                color="neutral"
                variant="ghost"
                icon="i-lucide-x"
                @click="open = false"
              />
            </div>
            <div>
            <h2 class="mb-1">Từ khóa</h2>
            <UInput v-model="filter.keyword"class="mb-6" icon="i-lucide-search" type="search" placeholder="Nhập từ khóa..."/>
            </div>
              <div class="">
              <h2 class="mb-2">Trạng thái</h2>
              <UCheckbox v-model="filter.completed" label="Đã đánh dấu hoàn thành" size="xl" font="xl" class="mb-2"></UCheckbox>
              <UCheckbox v-model="filter.incompleted" label="Không được đánh dấu hoàn thành" size="xl" font="xl" class="mb-4"></UCheckbox>
              <h2 class="mb-2">Ngày hết hạn</h2>
              <UCheckbox v-model="filter.noDueDate" label="Không có ngày hết hạn" size="xl" font="xl" class="mb-2"></UCheckbox>
              <UCheckbox v-model="filter.dueTomorrow" label="Sẽ hết hạn vào ngày mai" size="xl" font="xl" class="mb-2"></UCheckbox>
              <UCheckbox v-model="filter.overdue" label="Quá hạn" size="xl" font="xl"></UCheckbox>
              <!-- <UCheckbox v-model="filter." label="Hết hạn vào tuần sau" size="xl" font="xl"></UCheckbox>
              <UCheckbox v-model="filter." label="Hết hạn vào tháng sau" size="xl" font="xl"></UCheckbox> -->
              </div>

            <Placeholder class="size-full min-h-48" />
          </template>
        </UPopover>
        <UAvatar
          :src="user?.avatar?.startsWith('http') ? user?.avatar : undefined"
          :alt="user?.name"
          :ui="{ rounded: 'full' }"
          class="cursor-pointer ms-5"
          size="2xl"
        >
          <template #fallback>
            <span class="text-xs font-bold uppercase">{{
              getInitial(user?.name)
            }}</span>
          </template>
        </UAvatar>
        <UButton color="red" @click="logout">Đăng xuất</UButton>
      </div>
    </nav>

    <!-- Board content -->
    <div class="p-4">
      <div class="flex gap-4 overflow-x-auto">
        <UCard
          v-for="list in lists"
          :key="list.id"
          class="w-68 min-h-[300px] flex-shrink-0 bg-blue-50 shadow p-4 flex flex-col justify-between"
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
            @change="(event) => onCardDrop(event, list.id)"
            class="flex flex-col gap-2 mb-2"
          >
            <template #item="{ element }">
              <div
                v-if="checkVisible(element)"
                class="bg-white p-3 rounded shadow cursor-pointer space-y-2"
              >
                <!-- Title & Actions -->
                <div class="flex items-center gap-2 group w-full">
                  <UCheckbox
                    :model-value="element.completed"
                    @update:model-value="(val) => toggleCheckcard(element, val)"
                    size="xs"
                  />
                  <div class="flex-1">
                    <div v-if="editingCardId !== element.ID" @click="openCard(element)">
                      <div class="font-medium truncate cursor-pointer">
                        {{ element.title }}
                      </div>
                    </div>
                    <div v-else>
                      <UInput
                        v-model="editingTitle"
                        @blur="saveTitle(element)"
                        @keydown.enter.prevent="saveTitle(element)"
                        class="w-full text-sm"
                        autofocus
                      />
                    </div>
                  </div>

                  <div
                    class="flex gap-1 opacity-0 group-hover:opacity-100 transition"
                  >
                    <UButton
                      icon="i-lucide-pencil"
                      size="xs"
                      variant="ghost"
                      @click.stop="startEditing(element)"
                    />
                    <UButton
                      icon="i-lucide-x"
                      size="xs"
                      variant="ghost"
                      @click.stop="deleteCard(element)"
                    />
                  </div>
                </div>

                <!-- Dates -->
                <div class="flex items-center">
                  <div
                    class="text-xs text-gray-500 flex items-center gap-1 me-1.5"
                  >
                    <span
                    class="flex items-center"
                      v-if="element.start_date && element.end_date"
                      :class="`px-3 rounded text-sm  text-black ${getDateStatusColor(
                        element
                      )}`"
                    >
                    <UIcon name="i-heroicons-clock" class="size-4 me-0.5"/>
                      {{ formatDateOnCard(element.start_date) }} |
                      {{ formatDateOnCard(element.end_date) }}
                    </span>
                  </div>

                  <UTooltip
                    v-if="element.description"
                    :delay-duration="0"
                    text="Thẻ đã có mô tả"
                  >
                    <UIcon name="i-lucide-file-text" class="text-gray-600" />
                  </UTooltip>
                </div>

                <!-- Checklist count -->
                <div class="text-xs text-gray-500 flex items-center gap-2">
                  <UIcon name="i-lucide-check-square" class="text-gray-400" />
                  <span>
                    {{
                      element.checklist_items?.filter((i) => i.completed)
                        .length || 0
                    }}
                    /
                    {{ element.checklist_items?.length || 0 }}
                  </span>

                  <!-- Comment count -->
                  <UIcon
                    name="i-lucide-message-circle"
                    class="text-gray-400 ms-2"
                  />
                  <span>{{ element.comments?.length || 0 }}</span>
                </div>

                <!-- Members avatar -->
                <UAvatarGroup :max="3" size="2xs" class="pe-2">
                  <UAvatar
                    v-for="member in element.members || []"
                    :key="member.id"
                    :src="member.user.avatar"
                    size="2xs"
                  />
                </UAvatarGroup>
              </div>
            </template>
          </draggable>

          <!-- Thêm thẻ -->
          <div v-if="list.showAddCard">
            <UInput v-model="list.newCardTitle" class="mb-2" />
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
    :close="{ color: 'gray', class: 'rounded-full' }"
    class="max-w-5xl"
  >
    <template #body>
      <div class="flex flex-col md:flex-row gap-6">
        <!-- Bên trái - Thông tin thẻ -->
        <div class="flex-1 space-y-4 w-500">
          <!-- Thành viên, nhãn, ngày -->
          <div class="flex flex-wrap items-center gap-3">
            <!-- Thành viên -->
            <div>
              <h3 class="text-base font-semibold mb-3">Thành viên</h3>
              <div class="flex align-center justify-between w-150">
                <div class="flex align-center">
                  <UAvatarGroup :max="4" size="xl" class="pe-2">
                    <UAvatar
                      v-for="member in members"
                      :key="member.id"
                      :src="
                        member.user.avatar?.startsWith('http')
                          ? member.user.avatar
                          : undefined
                      "
                      :alt="member.user.name"
                      :ui="{ rounded: 'full' }"
                      class="cursor-pointer"
                    >
                      <template #fallback>
                        <span class="text-xs font-bold uppercase">{{
                          getInitial(member.user.name)
                        }}</span>
                      </template>
                    </UAvatar>
                  </UAvatarGroup>

                  <UPopover
                    class="rounded-full"
                    title="Thành viên"
                    :close="{
                      color: 'primary',
                      variant: 'outline',
                      class: 'rounded-full',
                    }"
                  >
                    <UButton
                      icon="i-lucide-plus"
                      color="neutral"
                      variant="subtle"
                      class="rounded-full"
                      size="xl"
                    />

                    <template #content>
                      <h4 class="text-center pt-3"><b>Thành viên</b></h4>
                      <div class="p-3">
                        <div class="flex gap-2 items-center mt-2 mb-3">
                          <USelectMenu
                            v-model="selectedUserId"
                            :items="allUsersOptions"
                            class="flex-1 w-48 h-8"
                          />
                          <UButton
                            color="primary"
                            variant="soft"
                            :disabled="!selectedUserId"
                            @click="addMember"
                          >
                            Thêm
                          </UButton>
                        </div>
                        <div
                          v-for="member in members"
                          :key="member.id"
                          class="flex items-center gap-3"
                        >
                          <UAvatar
                            :src="
                              member.user.avatar?.startsWith('http')
                                ? member.user.avatar
                                : undefined
                            "
                            :alt="member.user.name"
                            :ui="{ rounded: 'full' }"
                            class="cursor-pointer"
                          >
                            <template #fallback>
                              <span class="text-xs font-bold uppercase">{{
                                getInitial(member.user.name)
                              }}</span>
                            </template>
                          </UAvatar>
                          <div class="text-sm font-medium p-2">
                            {{ member.user.email }}
                          </div>

                          <UButton
                            color="red"
                            size="xs"
                            variant="ghost"
                            icon="i-lucide-x"
                            class="ml-auto"
                            @click="removeMember(member)"
                          />
                        </div>
                      </div>
                    </template>
                  </UPopover>
                </div>

                <!-- Ngày -->
                <UButton
                  icon="i-lucide-calendar"
                  @click="showDateModal = true"
                  class="flex items-center gap-2 border-1"
                  color="none"
                >
                  {{ formattedDates }}

                  <template v-if="dateStatus === 'completed'">
                    <span
                      class="px-2 py-0.5 bg-green-500 text-white text-xs rounded"
                    >
                      Hoàn tất
                    </span>
                  </template>
                  <template v-else-if="dateStatus === 'soon'">
                    <span
                      class="px-2 py-0.5 bg-yellow-400 text-white text-xs rounded"
                    >
                      Sắp tới hạn
                    </span>
                  </template>
                  <template v-else-if="dateStatus === 'overdue'">
                    <span
                      class="px-2 py-0.5 bg-red-500 text-white text-xs rounded"
                    >
                      Quá hạn
                    </span>
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
                        <UButton @click="confirmDates" color="primary"
                          >Lưu</UButton
                        >
                        <UButton @click="removeDates" color="gray"
                          >Gỡ bỏ</UButton
                        >
                      </div>
                    </div>
                  </template>
                </UModal>
              </div>
            </div>

            <!-- <div class="flex items-center gap-2">
              <span class="font-semibold">Nhãn:</span>
              <span class="w-6 h-6 bg-green-500 rounded" title="a"></span>
              <span class="w-6 h-6 bg-yellow-400 rounded" title="b"></span>
              <UButton
                size="xs"
                icon="i-lucide-plus"
                color="gray"
                variant="soft"
              />
            </div> -->
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
          <h3 class="font-semibold pb-3">Nhận xét</h3>
          <div
            v-for="comment in comments"
            :key="comment.id"
            class="text-sm mb-2"
          >
            <div class="flex items-center gap-1 mb-1">
              <!-- <div
                class="rounded-full bg-blue-600 text-white w-8 h-8 flex items-center justify-center font-bold"
              >
                {{ getInitial(comment.User.name) }}
              </div> -->
              <UAvatar
                :src="
                  comment.User.avatar?.startsWith('http')
                    ? comment.User.avatar
                    : undefined
                "
                :alt="comment.User.name"
                :ui="{ rounded: 'full' }"
                class="cursor-pointer"
                size="xl"
              >
                <template #fallback>
                  <span class="text-xs font-bold uppercase">{{
                    getInitial(comment.User.name)
                  }}</span>
                </template>
              </UAvatar>
              <div>
                <div class="flex">
                  <span class="font-semibold">{{ comment.User.name }}:</span>
                  <p class="text-black ps-1">{{ comment.content }}</p>
                </div>

                <span class="text-gray-500 font-mono">{{
                  formatDate(comment.CreatedAt)
                }}</span>
              </div>
            </div>
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

const open = ref(false);
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
    UserID: user.value?.id,
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

async function toggleCheckcard(element, newValue) {
  element.completed = newValue;
  await $fetch(`http://localhost:3001/api/cards/${element.ID}`, {
    method: "PUT",
    body: { completed: element.completed },
  });
}

async function onCardDrop(event, toListId) {
  const { added, moved, removed } = event;
  const toList = lists.value.find((list) => list.id === toListId);
  if (!toList) return;

  await Promise.all(
    toList.cards.map((card, index) =>
      $fetch(`http://localhost:3001/api/cards/${card.ID}`, {
        method: "PUT",
        body: {
          list_id: toListId,
          position: index,
        },
      })
    )
  );
}
const editingCardId = ref(null);
const editingTitle = ref("");

function startEditing(card) {
  editingCardId.value = card.ID;
  editingTitle.value = card.title;
}

async function saveTitle(card) {
  if (!editingTitle.value.trim() || editingTitle.value === card.title) {
    editingCardId.value = null;
    return;
  }
  await $fetch(`http://localhost:3001/api/cards/${card.ID}`, {
    method: "PUT",
    body: {
      title: editingTitle.value,
    },
  });
  card.title = editingTitle.value;
  editingCardId.value = null;
}
async function deleteCard(element) {
  if (!confirm("Bạn có chắc muốn xoá thẻ này không?")) return;

  await $fetch(`http://localhost:3001/api/cards/${element.ID}`, {
    method: "DELETE",
  });
  const listTarget = lists.value.find((l) => l.id === element.list_id);
  if (listTarget) {
    listTarget.cards = listTarget.cards.filter((c) => c.ID !== element.ID);
  }
}

// --- Modal open card---
function openCard(card) {
  selectedCard.value = card;
  showCardModal.value = true;
  fetchMembers();
  fetchAllUsers();
  fetchComments(card.ID);
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
// avatar
function getInitial(name) {
  if (!name) return "?";
  return name.trim().charAt(0).toUpperCase();
}
function formatDate(dateStr) {
  const date = new Date(dateStr);
  return date.toLocaleString("vi-VN", {
    hour: "2-digit",
    minute: "2-digit",
    day: "2-digit",
    month: "2-digit",
    year: "numeric",
  });
}
function formatDateOnCard(dateStr) {
  const date = new Date(dateStr);
  return date.toLocaleString("vi-VN", {
    day: "2-digit",
    month: "2-digit",
  });
}
async function addComment() {
  if (!newComment.value.trim()) return;
  await $fetch(
    `http://localhost:3001/api/cards/${selectedCard.value.ID}/comments`,
    {
      method: "POST",
      body: {
        content: newComment.value,
        UserID: user.value?.id,
      },
    }
  );
  newComment.value = "";
  fetchComments(selectedCard.value.ID);
}

// --- Checklist ---
async function fetchChecklist() {
  const { data } = await useFetch(
    `http://localhost:3001/api/cards/${selectedCard.value.ID}/checklist`
  );
  checklist.value = data.value || [];
}

async function addChecklist() {
  if (!newChecklist.value.trim()) return;
  const res = await $fetch(
    `http://localhost:3001/api/cards/${selectedCard.value.ID}/checklist`,
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
  await $fetch(`http://localhost:3001/api/cards/${selectedCard.value.ID}`, {
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
  await $fetch(`http://localhost:3001/api/cards/${selectedCard.value.ID}`, {
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
  await $fetch(`http://localhost:3001/api/cards/${selectedCard.value.ID}`, {
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
  const endTimeStr = end.toLocaleTimeString("vi-VN", {
    hour: "2-digit",
    minute: "2-digit",
  });

  const endDateStr = end.toLocaleDateString("vi-VN", {
    day: "2-digit",
    month: "2-digit",
  });

  return `${startDateStr} || ${endTimeStr}  ${endDateStr}`;
});

const dateStatus = computed(() => {
  if (!selectedCard.value) return "normal";
  if (selectedCard.value.completed) return "completed";
  if (!selectedCard.value.end_date) return "normal";

  const now = new Date();
  const end = new Date(selectedCard.value.end_date);

  if (now > end) return "overdue";

  const diff = (end.getTime() - now.getTime()) / (1000 * 60 * 60 * 24);
  if (diff <= 1) return "soon";

  return "normal";
});
function getDateStatusColor(card) {
  if (!card) return "bg-gray-100";
  if (card.completed) return "bg-green-400";
  if (!card.end_date) return "bg-gray-100";

  const now = new Date();
  const end = new Date(card.end_date);

  if (now > end) return "bg-red-400";

  const diff = (end.getTime() - now.getTime()) / (1000 * 60 * 60 * 24);
  if (diff <= 1) return "bg-yellow-200";

  return "bg-gray-100";
}

const members = ref([]);
const allUsers = ref([]);
const selectedUserId = ref("");

async function fetchMembers() {
  if (!selectedCard.value?.ID) return;
  const { data } = await useFetch(
    `http://localhost:3001/api/cards/${selectedCard.value.ID}/members`
  );
  members.value = data.value || [];
}

async function fetchAllUsers() {
  const { data } = await useFetch("http://localhost:3001/api/users");
  allUsers.value = data.value || [];
}
const allUsersOptions = computed(() =>
  allUsers.value.map((user) => ({
    label: user.email,
    value: user.ID,
  }))
);
async function addMember() {
  if (!selectedUserId.value || !selectedCard.value?.ID) return;
  await $fetch(
    `http://localhost:3001/api/cards/${selectedCard.value.ID}/members`,
    {
      method: "POST",
      body: {
        user_id: Number(selectedUserId.value.value),
        role: "viewer",
      },
    }
  );
  selectedUserId.value = null;
  await fetchMembers();
}

async function removeMember(member) {
  if (!selectedCard.value?.ID || !member?.user?.ID) return;
  await $fetch(
    `http://localhost:3001/api/cards/${selectedCard.value.ID}/members/${member.user.ID}`,
    {
      method: "DELETE",
    }
  );
  await fetchMembers();
}
const props = defineProps(["card"]);


const filter = reactive({
  keyword: '',
  completed: false,
  incompleted: false,
  noDueDate: false,
  dueTomorrow: false,
  overdue: false
});

function checkVisible(card) {
  if (filter.keyword && !card.title.toLowerCase().includes(filter.keyword.toLowerCase())) {
    return false;
  }

  if (filter.completed  && !card.completed) return false;
if (filter.incompleted  && card.completed) return false;

  const now = new Date();
  const end = card.end_date ? new Date(card.end_date) : null;
  const diff = end ? (end.getTime() - now.getTime()) / (1000 * 60 * 60 * 24) : null;

  const matchDueFilters = [];

  if (filter.noDueDate) {
    matchDueFilters.push(!end);
  }
  if (filter.dueTomorrow) {
    matchDueFilters.push(end && diff <= 1 && diff > 0);
  }
  if (filter.overdue) {
    matchDueFilters.push(end && now > end && !card.completed);
  }

  if (
    filter.noDueDate ||
    filter.dueTomorrow ||
    filter.overdue
  ) {
    // Nếu có tick ít nhất 1 filter ngày, phải match ít nhất 1
    if (!matchDueFilters.some(Boolean)) return false;
  }

  return true;
}
watch(
  () => filter.completed,
  (val) => {
    if (val) filter.incompleted = false
  }
)

watch(
  () => filter.incompleted,
  (val) => {
    if (val) filter.completed = false
  }
)


</script>
