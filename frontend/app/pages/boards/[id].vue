<template>
  <div class="min-h-screen bg-gray-100">
    <!-- Navbar -->
    <nav class="flex justify-between items-center px-6 py-4 bg-white shadow-sm">
      <h1 class="text-2xl font-bold text-primary">{{ board?.title || "Board" }}</h1>
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
                <button class="text-xs text-blue-500" @click="list.editing = true">Sửa</button>
                <button class="text-xs text-red-500" @click="deleteList(list.id)">Xóa</button>
              </div>
            </div>
            <div v-else class="flex gap-2 w-full">
              <UInput v-model="list.title" size="xs" class="w-full" />
              <UButton size="xs" @click="updateList(list)">Lưu</UButton>
              <UButton size="xs" color="gray" @click="list.editing = false">Hủy</UButton>
            </div>
          </div>

          <!-- Thẻ (cards) -->
          <draggable
            v-model="list.cards"
            group="cards"
            item-key="id"
            :data-list-id="list.id"
            @end="event => onCardDrop(event, list.id)"
            class="flex flex-col gap-2 mb-2"
          >
            <template #item="{ element }">
              <div class="bg-gray-100 p-2 rounded shadow cursor-pointer" @click="openCard(element)">
                {{ element.title }}
              </div>
            </template>
          </draggable>  

          <!-- Thêm thẻ -->
          <div v-if="list.showAddCard">
            <UInput v-model="list.newCardTitle" placeholder="Tên thẻ..." class="mb-2" />
            <div class="flex gap-2">
              <UButton size="xs" @click="createCard(list)">Thêm</UButton>
              <UButton size="xs" color="gray" @click="list.showAddCard = false">Hủy</UButton>
            </div>
          </div>
          <div v-else>
            <button class="text-sm text-gray-500 hover:text-primary" @click="list.showAddCard = true">
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
            <UInput v-model="newListTitle" placeholder="Tên cột..." class="mb-2" />
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
      <!-- Bên trái: Thông tin thẻ -->
      <div class="flex-1 space-y-4">

        <!-- Thành viên, nhãn, ngày -->
        <div class="flex flex-wrap items-center gap-3">
          <div class="flex items-center gap-2">
            <span class="font-semibold">Thành viên:</span>
            <div class="rounded-full bg-blue-600 text-white w-8 h-8 flex items-center justify-center font-bold">
              MH
            </div>
            <UButton size="xs" icon="i-lucide-plus" color="gray" variant="soft" />
          </div>

          <div class="flex items-center gap-2">
            <span class="font-semibold">Nhãn:</span>
            <span class="w-6 h-6 bg-green-500 rounded" title="a"></span>
            <span class="w-6 h-6 bg-yellow-400 rounded" title="b"></span>
            <UButton size="xs" icon="i-lucide-plus" color="gray" variant="soft" />
          </div>

          <div class="flex items-center gap-2">
            <span class="font-semibold">Ngày:</span>
            <UBadge color="gray" variant="soft">
              {{ selectedCard?.start_date }} - {{ selectedCard?.end_date }}
            </UBadge>
            <UBadge color="green" variant="solid">Hoàn tất</UBadge>
          </div>
        </div>

        <!-- Mô tả -->
        <div>
          <h3 class="font-semibold flex items-center gap-2 mt-4">
            <i class="i-lucide-align-left" /> Mô tả
          </h3>
          <p class="text-gray-700 mt-1">{{ selectedCard?.description || "Không có mô tả" }}</p>
        </div>

        <!-- Checklist -->
<!-- Checklist -->
<div>
  <h3 class="font-semibold flex items-center gap-2 mt-4">
    <i class="i-lucide-check-square" /> Việc cần làm
  </h3>

  <div v-for="item in checklist" :key="item.id" class="flex items-center gap-2 mt-2">
    <UCheckbox v-model="item.completed" @change="toggleChecklist(item)" />
    <span :class="{ 'line-through text-gray-500': item.completed }">{{ item.content }}</span>
    <UButton icon="i-lucide-trash" color="red" variant="ghost" size="xs" @click="deleteChecklist(item.id)" />
  </div>

  <div class="flex gap-2 mt-2">
    <UInput v-model="newChecklist" placeholder="Thêm mục mới..." class="w-full" />
    <UButton size="xs" @click="addChecklist">Thêm</UButton>
  </div>
</div>


      </div>
      <!-- Bình luận -->
<div>
  <h3 class="font-semibold">Nhận xét</h3>
  <div v-for="comment in comments" :key="comment.id" class="text-sm mb-2">
    <span class="text-gray-500">{{ comment.created_at }}</span><br />
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
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import draggable from 'vuedraggable'

const route = useRoute()
const router = useRouter()

const boardId = route.params.id
const user = ref(null)
const board = ref({})
const lists = ref([])

const newListTitle = ref("")
const showCreateInput = ref(false)

// Đăng xuất
function logout() {
  localStorage.clear()
  router.push("/login")
}

// Tải dữ liệu
onMounted(async () => {
  const saved = localStorage.getItem("user")
  if (!saved) return router.push("/login")
  user.value = JSON.parse(saved)

  const { data: boardData } = await useFetch(`http://localhost:3001/api/boards/${boardId}`)
  const { data: listData } = await useFetch(`http://localhost:3001/api/boards/${boardId}/lists`)

  board.value = boardData.value
  lists.value = listData.value.map(list => ({
    ...list,
    showAddCard: false,
    newCardTitle: "",
    editing: false,
    cards: (list.cards || []).sort((a, b) => a.position - b.position)
  }))
})

// Tạo cột mới
async function createList() {
  if (!newListTitle.value.trim()) return
  const { data, error } = await useFetch("http://localhost:3001/api/lists", {
    method: "POST",
    body: {
      title: newListTitle.value,
      board_id: parseInt(boardId)
    }
  })

  if (error.value) {
    alert("Tạo cột thất bại")
    return
  }

  lists.value.push({
    ...data.value,
    showAddCard: false,
    newCardTitle: "",
    editing: false,
    cards: []
  })

  newListTitle.value = ""
  showCreateInput.value = false
}
// Sửa tiêu đề cột
async function updateList(list) {
  try {
    await $fetch(`http://localhost:3001/api/lists/${list.id}`, {
      method: "PUT",
      body: { title: list.title }
    })
    list.editing = false
  } catch (err) {
    alert("Không thể cập nhật cột")
  }
}

// Xoá cột
async function deleteList(id) {
  if (!confirm("Bạn có chắc muốn xoá cột này không?")) return
  try {
    await $fetch(`http://localhost:3001/api/lists/${id}`, {
      method: "DELETE"
    })
    lists.value = lists.value.filter(l => l.id !== id)
  } catch (err) {
    alert("Không thể xoá cột")
  }
}
// Tạo thẻ
async function createCard(list) {
  const payload = {
    title: list.newCardTitle,
    list_id: list.id,
    position: list.cards.length,
    description: "",
    start_date: "",
    end_date: "",
    user_id: user.value?.id || 1,
  }

  const { data, error } = await useFetch("http://localhost:3001/api/cards", {
    method: "POST",
    body: payload
  })

  if (error.value) {
    console.error("Lỗi khi tạo card:", error.value)
    alert("Không thể tạo thẻ.")
    return
  }

  list.cards.push(data.value)
  list.newCardTitle = ""
  list.showAddCard = false
}

// Kéo thả thẻ
async function onCardDrop(event, toListId) {
  const movedCard = event.item?._underlying_vm_

  if (!movedCard?.id) {
    console.error("Card không có ID khi drop:", movedCard)
    return
  }

  const toList = lists.value.find(list => list.id === toListId)
  const newCards = toList.cards

  for (let i = 0; i < newCards.length; i++) {
    const card = newCards[i]

    await $fetch(`http://localhost:3001/api/cards/${card.id}`, {
      method: "PUT",
      body: {
        list_id: toListId,
        position: i
      }
    })
  }
}

  const showCardModal = ref(false)
const selectedCard = ref(null)

function openCard(card) {
  selectedCard.value = card
  showCardModal.value = true
  fetchComments(card.id)
  fetchChecklist()

}


// Script setup trong modal card
const comments = ref([])
const newComment = ref("")

async function fetchComments(cardId) {
  const { data } = await useFetch(`http://localhost:3001/api/cards/${cardId}/comments`)
  // console.log("Comments:", data.value)
  comments.value = data.value || []
}

async function addComment() {
  if (!newComment.value.trim()) return

  await $fetch(`http://localhost:3001/api/cards/${selectedCard.value.id}/comments`, {
    method: "POST",
    body: {
      content: newComment.value,
      user_id: user.value?.id || 1
    }
  })

  newComment.value = ""
  fetchComments(selectedCard.id)
}

// Gọi fetch khi mở modal
watch(() => showCardModal.value, (open) => {
  if (open && selectedCard.value?.id) {
    fetchComments(selectedCard.value.id)
  }
})

const checklist = ref([])
const newChecklist = ref("")

async function fetchChecklist() {
  const res = await $fetch(`http://localhost:3001/api/cards/${selectedCard.value.id}/checklist`)
  checklist.value = res
}

async function addChecklist() {
  if (!newChecklist.value.trim()) return
  const res = await $fetch(`http://localhost:3001/api/cards/${selectedCard.value.id}/checklist`, {
    method: "POST",
    body: {
      content: newChecklist.value,
      completed: false,
    },
  })
  checklist.value.push(res)
  newChecklist.value = ""
}

async function toggleChecklist(item) {
  await $fetch(`http://localhost:3001/api/checklist/${item.id}`, {
    method: "PUT",
    body: {
      content: item.content,
      completed: item.completed,
    },
  })
}

async function deleteChecklist(id) {
  await $fetch(`http://localhost:3001/api/checklist/${id}`, {
    method: "DELETE",
  })
  checklist.value = checklist.value.filter(i => i.id !== id)
}



</script>
