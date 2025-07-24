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
              <h2 class="text-highlighted font-semibold">Bộ lọc</h2>
              <UButton
                color="neutral"
                variant="ghost"
                icon="i-lucide-x"
                @click="open = false"
              />
            </div>
            <div>
              <h2 class="mb-1">Từ khóa</h2>
              <UInput
                v-model="filter.keyword"
                class="mb-6"
                icon="i-lucide-search"
                type="search"
                placeholder="Nhập từ khóa..."
              />
            </div>
            <div class="">
              <h2 class="mb-2">Thành viên</h2>
              <UCheckbox
                v-model="filter.noMembers"
                label="Không có thành viên"
                size="xl"
                font="xl"
                class="mb-2"
              />
              <UCheckbox
                v-model="filter.assignedToMe"
                label="Được chỉ định cho tôi"
                size="xl"
                font="xl"
                class="mb-2"
              />  

              <!-- <div class="flex items-center mb-1">
                <UCheckbox class="me-2.5" size="xl" />
                <USelectMenu
                  v-model="filter.memberIds"
                  :items="listuseroption"
                  multiple
                  value-attribute="value"
                  option-attribute="name"
                  placeholder="Chọn thành viên"
                >
                  <template #option="{ option, selected }">
                    <div class="flex items-center gap-3">
                      <UCheckbox :model-value="selected" class="pointer-events-none" />
                      <UAvatar :src="option.avatar" :alt="option.name" size="xs" />
                      <div class="flex flex-col">
                        <span class="text-sm font-medium">{{ option.name }}</span>
                        <span class="text-xs text-gray-500">{{ option.email }}</span>
                      </div>
                    </div>
                  </template>

                  <template #label>
                    <span class="text-sm">
                      {{ filter.memberIds.length === 0
                        ? 'Chọn thành viên'
                        : `${filter.memberIds.length} thành viên được chọn` }}
                    </span>
                  </template>
                </USelectMenu>
              </div> -->

              <p class="text-[13px] text-gray-600 mb-3">Tìm kiếm các thẻ, các thành viên, các nhãn và hơn thế nữa.</p>
              <h2 class="mb-2">Trạng thái</h2>
              <UCheckbox
                v-model="filter.completed"
                label="Đã đánh dấu hoàn thành"
                size="xl"
                font="xl"
                class="mb-2"
              ></UCheckbox>
              <UCheckbox
                v-model="filter.incompleted"
                label="Không được đánh dấu hoàn thành"
                size="xl"
                font="xl"
                class="mb-4"
              ></UCheckbox>
              <h2 class="mb-2">Ngày hết hạn</h2>
              <UCheckbox
                v-model="filter.noDueDate"
                label="Không có ngày hết hạn"
                size="xl"
                font="xl"
                class="mb-2"
              ></UCheckbox>
              <UCheckbox
                v-model="filter.overdue"
                label="Quá hạn"
                size="xl"
                font="xl"
                class="mb-2"
              ></UCheckbox>
              <UCheckbox
                v-model="filter.dueTomorrow"
                label="Sẽ hết hạn vào ngày mai"
                size="xl"
                font="xl"
                class="mb-2"
              ></UCheckbox>
              <UCheckbox
                v-model="filter.expireNextWeek"
                label="Hết hạn vào tuần sau"
                size="xl"
                font="xl"
                class="mb-2"
              ></UCheckbox>
              <UCheckbox
                v-model="filter.expireNextMonth"
                label="Hết hạn vào tháng sau"
                size="xl"
                font="xl"
                class="mb-4"
              ></UCheckbox>
            </div>
          </template>
        </UPopover>
        <UAvatar
          :src="user?.avatar?.startsWith('https') ? user?.avatar : undefined"
          :alt="user?.name"
          :ui="{ rounded: 'full' }"
          class="cursor-pointer ms-5"
          size="2xl"
        >
        </UAvatar>
        <UButton color="red" @click="logout">Đăng xuất</UButton>
      </div>
    </nav>

    <!-- Board content -->

    <div class="p-4 w-full">
      <div class="flex gap-4 overflow-x-auto">
        <!-- Kéo thả cột (list) -->
        <draggable
          v-model="lists"
          item-key="id"
          direction="horizontal"
          class="flex gap-4"
          @change="onListDrop"
        >
          <template #item="{ element }">
            <UCard
              class="w-68 min-h-[300px] flex-shrink-0 bg-blue-50 shadow p-4 flex flex-col justify-between"
            >
              <div class="flex justify-between items-center mb-2 group">
                <div
                  v-if="editingListId !== element.id"
                  class="flex justify-between w-full"
                >
                  <h3 class="font-extrabold text-xl truncate cursor-pointer break-words">
                    {{ element.title }}
                  </h3>
                  <div class="flex gap-1 opacity-0 group-hover:opacity-100">
                    <UButton
                      icon="i-lucide-pencil"
                      color="info"
                      size="xs"
                      variant="ghost"
                      @click.stop="startEditingList(element)"
                    />
                    <UButton
                      icon="i-lucide-trash"
                      color="error"
                      size="xs"
                      variant="ghost"
                      @click.stop="deleteList(element.id)"
                    />
                  </div>
                </div>
                <div v-else class="flex gap-2 w-full">
                  <UInput
                    v-model="editingListTitle"
                    size="xs"
                    class="w-full"
                    @blur="saveList(element)"
                    @keydown.enter.prevent="saveList(element)"
                    @keydown.esc.prevent="cancelEditList"
                    autofocus
                  />
                </div>
              </div>

              <!-- Kéo thả card bên trong -->
              <draggable
                v-model="element.cards"
                group="cards"
                item-key="id"
                :data-list-id="element.id"
                @change="(event) => onCardDrop(event, element.id)"
                class="flex flex-col gap-2 mb-2"
              >
                <template #item="{ element: card }">
                  <div
                    v-if="checkVisible(card)"
                    class="bg-white p-3 rounded-2xl shadow cursor-pointer space-y-2 w-full"
                  >
                    <!-- Checkbox & title -->
                    <div class="flex items-center gap-2 group w-full">
                      <UTooltip
                        :text="
                          card.completed
                            ? 'Đánh dấu chưa hoàn tất'
                            : 'Đánh dấu hoàn tất'
                        "
                        :delay-duration="0"
                      >
                        <UCheckbox
                          :model-value="card.completed"
                          @update:model-value="
                            (val) => toggleCheckcard(card, val)
                          "
                          size="sm"
                          class=""
                        />
                      </UTooltip>

                      <div class="flex-1">
                        <div
                          v-if="editingCardId !== card.ID"
                          @click="openCard(card)"
                        >
                          <div class="font-medium truncate cursor-pointer">
                            {{ card.title }}
                          </div>
                        </div>
                        <div v-else>
                          <UInput
                            v-model="editingTitle"
                            @blur="saveTitle(card)"
                            @keydown.enter.prevent="saveTitle(card)"
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
                          color="info"
                          size="xs"
                          variant="ghost"
                          @click.stop="startEditing(card)"
                        />
                        <UButton
                          icon="i-lucide-trash"
                          color="error"
                          size="xs"
                          variant="ghost"
                          @click.stop="deleteCard(card)"
                        />
                      </div>
                    </div>

                    <!-- Dates -->
                    <div class="flex items-center" @click="openCard(card)">
                      <UTooltip :text="tooltipDate(card)">
                        <div
                          class="text-xs text-gray-500 flex items-center gap-1 me-1.5"
                        >
                          <span
                            v-if="card.start_date && card.end_date"
                            :class="`flex items-center px-3 rounded text-sm text-black ${getDateStatusColor(
                              card
                            )}`"
                          >
                            <UIcon
                              name="i-heroicons-clock"
                              class="size-4 me-0.5"
                            />
                            {{ formatDateOnCard(card.start_date) }} |
                            {{ formatDateOnCard(card.end_date) }}
                          </span>
                        </div>
                      </UTooltip>

                      <UTooltip
                        v-if="card.description"
                        :delay-duration="0"
                        text="Thẻ đã có mô tả"
                      >
                        <UIcon
                          name="i-lucide-file-text"
                          class="text-gray-600"
                        />
                      </UTooltip>
                    </div>

                    <!-- Checklist + comment count -->
                    <div class="text-xs text-gray-500 flex items-center gap-2" @click="openCard(card)">
                      <UTooltip
                        v-if="card.checklist_items?.length"
                        text="Mục trong danh sách công việc"
                      >
                        <div class="flex items-center gap-1">
                          <UIcon
                            name="i-lucide-check-square"
                            :class="{
                              'text-green-500':
                                card.checklist_items?.length > 0 &&
                                card.checklist_items?.filter((i) => i.completed)
                                  .length === card.checklist_items?.length,
                              'text-gray-400': !(
                                card.checklist_items?.length > 0 &&
                                card.checklist_items?.filter((i) => i.completed)
                                  .length === card.checklist_items?.length
                              ),
                            }"
                          />
                          <span>
                            {{
                              card.checklist_items?.filter((i) => i.completed)
                                .length || 0
                            }}/
                            {{ card.checklist_items?.length || 0 }}
                          </span>
                        </div>
                      </UTooltip>

                      <UTooltip v-if="card.comments?.length" text="Bình luận">
                        <div class="flex items-center gap-1">
                          <UIcon
                            name="i-lucide-message-circle"
                            class="text-gray-400 ms-2"
                          />
                          <span>{{ card.comments?.length || 0 }}</span>
                        </div>
                      </UTooltip>
                    </div>

                    <!-- Members -->
                    <UTooltip v-if="card.members?.length" text="Thành viên" @click="openCard(card)">
                      <div class="flex justify-end">
                        <UAvatarGroup :max="3" size="4xs" class="pe-2">
                          <UAvatar
                            v-for="member in card.members || []"
                            :key="member.id"
                            :src="member.user.avatar"
                            size="2xs"
                          />
                        </UAvatarGroup>
                      </div>
                    </UTooltip>
                  </div>
                </template>
              </draggable>

              <!-- Thêm thẻ -->
              <div v-if="element.showAddCard">
                <UInput v-model="element.newCardTitle" class="mb-2" />
                <div class="flex gap-2">
                  <UButton size="xs" @click="createCard(element)">Thêm</UButton>
                  <UButton
                    size="xs"
                    color="gray"
                    @click="element.showAddCard = false"
                    >Hủy</UButton
                  >
                </div>
              </div>
              <div v-else>
                <button
                  class="text-sm text-gray-500 hover:text-primary"
                  @click="element.showAddCard = true"
                >
                  + Thêm thẻ
                </button>
              </div>
            </UCard>
          </template>
        </draggable>

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

  <!-- showCardModal -->
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
                  <UAvatarGroup :max="5" size="xl" class="pe-2">
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
              <div
                v-if="editingChecklistId !== item.id"
                class="flex items-center gap-2 flex-1"
              >
                <span
                  :class="{ 'line-through text-gray-500': item.completed }"
                  @click="startEditingChecklist(item)"
                  class="cursor-pointer w-full"
                >
                  {{ item.content }}
                </span>
                <UButton
                  size="xs"
                  icon="i-lucide-pencil"
                  @click="startEditingChecklist(item)"
                  color="none"
                />
              </div>
              <div v-else class="flex-1">
                <UInput
                  v-model="editingChecklistContent"
                  size="xs"
                  class="w-full"
                  @blur="saveChecklistContent(item)"
                  @keydown.enter.prevent="saveChecklistContent(item)"
                  @keydown.esc.prevent="cancelEditChecklist"
                  autofocus
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
        <div class="md:w-2/4 w-full pl-4">
          <div class="flex items-center pb-3">
            <UIcon name="i-lucide-message-circle" />
            <h3 class="font-semibold ms-1.5">Nhận xét hoạt động</h3>
          </div>
          <div
            v-for="comment in comments"
            :key="comment.ID"
            class="text-sm mb-2"
          >
            <div class="flex items-center gap-1 mb-1">
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
              />
              <div class="flex items-center w-full">
                <span class="font-bold text-black">{{
                  comment.User.name
                }}</span>
                <div class="items-center gap-1 ms-5">
                  <span class="text-blue-500 text-xs underline font-mono">{{
                    formatDate(comment.CreatedAt)
                  }}</span>
                  <span
                    v-if="isEdited(comment)"
                    class="italic text-gray-400 text-[11px] ms-1"
                    >(đã sửa)</span
                  >
                </div>
              </div>
            </div>

            <!-- Nội dung bl -->
            <div>
              <template v-if="editingCommentId === comment.ID">
                <UInput
                  v-model="editedContent"
                  @keyup.enter="saveComment(comment.ID)"
                  size="xs"
                  class="w-full ms-8"
                />
              </template>
              <template v-else>
                <div class="bg-old-neutral-100 ms-8 rounded-xl p-2 max-w-[260px]">
                  <p class="text-black break-words">{{ comment.content }}</p>
                </div>

                <div
                  v-if="comment.UserID === user.id"
                  class="flex gap-1 text-xs text-gray-500 ms-10 mt-1"
                >
                  <span
                    class="cursor-pointer hover:underline me-2"
                    @click="startEditingcmt(comment)"
                  >
                    Chỉnh sửa
                  </span>
                  <span
                    class="cursor-pointer hover:underline"
                    @click="deleteComment(comment.ID)"
                  >
                    Xóa
                  </span>
                </div>
              </template>
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
const currentUserId = ref(null);
// --- Fetch board + list ---
onMounted(async () => {
  const saved = localStorage.getItem("user");
  if (!saved) return router.push("/login");
  user.value = JSON.parse(saved);

  currentUserId.value = user.value.id;

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
    const position = lists.value.length > 0
    ? Math.max(...lists.value.map(l => l.position)) + 1
    : 0;
  const { data, error } = await useFetch("http://localhost:3001/api/lists", {
    method: "POST",
    body: {
      title: newListTitle.value,
      board_id: parseInt(boardId),
      position,
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

const editingListId = ref(null);
const editingListTitle = ref("");

function startEditingList(list) {
  editingListId.value = list.id;
  editingListTitle.value = list.title;
}

async function saveList(list) {
  if (!editingListTitle.value.trim() || editingListTitle.value === list.title) {
    editingListId.value = null;
    return;
  }
  await $fetch(`http://localhost:3001/api/lists/${list.id}`, {
    method: "PUT",
    body: { title: editingListTitle.value.trim() },
  });
  list.title = editingListTitle.value.trim();
  editingListId.value = null;
}

function cancelEditList() {
  editingListId.value = null;
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
async function onListDrop() {
  lists.value.forEach((list, index) => {
    list.position = index;
  });

  await Promise.all(
    lists.value.map((list) =>
      $fetch(`http://localhost:3001/api/lists/${list.id}`, {
        method: "PUT",
        body: {
          position: list.position,
          title: list.title,
        },
      })
    )
  );
}
// --- Card ---
async function createCard(list) {
  const title = list.newCardTitle.trim();
  if (!title){
    return;
  }
  const position =
    list.cards.length > 0 ? list.cards[list.cards.length - 1].position + 1 : 0;
  const payload = {
    title,
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
  await fetchComments(selectedCard.value.ID);
  const list = lists.value.find((l) => l.id === selectedCard.value.list_id);
  const card = list?.cards.find((c) => c.ID === selectedCard.value.ID);
  if (card) {
    card.comments = [...comments.value];
  }
}

const editingCommentId = ref(null);
const editedContent = ref("");

function startEditingcmt(comment) {
  editingCommentId.value = comment.ID;
  editedContent.value = comment.content;
}

async function saveComment(id) {
  if (!editedContent.value.trim()) return;
  await $fetch(
    `http://localhost:3001/api/cards/${selectedCard.value.ID}/comments/${id}`,
    {
      method: "PUT",
      body: {
        content: editedContent.value,
      },
    }
  );
  editingCommentId.value = null;
  editedContent.value = "";
  await fetchComments(selectedCard.value.ID);
}

async function deleteComment(id) {
  const ok = confirm("Bạn có chắc muốn xoá bình luận?");
  if (!ok) return;

  await $fetch(
    `http://localhost:3001/api/cards/${selectedCard.value.ID}/comments/${id}`,
    {
      method: "DELETE",
    }
  );
  await fetchComments(selectedCard.value.ID);
}
function isEdited(comment) {
  return (
    new Date(comment.UpdatedAt).getTime() -
      new Date(comment.CreatedAt).getTime() >
    1000
  );
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
  const list = lists.value.find((l) => l.id === selectedCard.value.list_id);
  const card = list?.cards.find((c) => c.ID === selectedCard.value.ID);
  if (card) {
    card.checklist_items = [...checklist.value];
  }
}

const editingChecklistId = ref(null);
const editingChecklistContent = ref("");

function startEditingChecklist(item) {
  editingChecklistId.value = item.id;
  editingChecklistContent.value = item.content;
}

async function saveChecklistContent(item) {
  if (
    !editingChecklistContent.value.trim() ||
    editingChecklistContent.value === item.content
  ) {
    editingChecklistId.value = null;
    return;
  }
  await $fetch(`http://localhost:3001/api/checklist/${item.id}`, {
    method: "PUT",
    body: {
      content: editingChecklistContent.value.trim(),
      completed: item.completed,
    },
  });
  item.content = editingChecklistContent.value.trim();
  editingChecklistId.value = null;
}

function cancelEditChecklist() {
  editingChecklistId.value = null;
}

async function toggleChecklist(item) {
  await $fetch(`http://localhost:3001/api/checklist/${item.id}`, {
    method: "PUT",
    body: {
      content: item.content,
      completed: item.completed,
    },
  });
  // await fetchChecklist();
  const list = lists.value.find((l) => l.id === selectedCard.value.list_id);
  const card = list?.cards.find((c) => c.ID === selectedCard.value.ID);
  if (card) {
    card.checklist_items = [...checklist.value];
  }
}

async function deleteChecklist(id) {
  if (!confirm("Bạn có chắc muốn xoá mục này không?")) return;
  await $fetch(`http://localhost:3001/api/checklist/${id}`, {
    method: "DELETE",
  });
  await fetchChecklist();
  const list = lists.value.find((l) => l.id === selectedCard.value.list_id);
  const card = list?.cards.find((c) => c.ID === selectedCard.value.ID);
  if (card) {
    card.checklist_items = [...checklist.value];
  }
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
watch(tempEndDate, (newEnd) => {
  if (!newEnd) return;
  if (!tempStartDate.value) return;

  const start = new Date(tempStartDate.value);
  const end = new Date(newEnd);

  if (end <= start) {
    const newStart = new Date(end);
    newStart.setDate(newStart.getDate() - 1);
    tempStartDate.value = newStart.toISOString().slice(0, 10);
  }
});
function formatDate(dateStr) {
  const date = new Date(dateStr);
  const now = new Date();
  const diffMs = now - date;

  const diffSec = Math.floor(diffMs / 1000);
  const diffMin = Math.floor(diffSec / 60);
  const diffHour = Math.floor(diffMin / 60);
  const diffDay = Math.floor(diffHour / 24);

  if (diffSec <= 60) {
    return "vừa xong";
  } else if (diffMin < 60) {
    return `${diffMin} phút trước`;
  } else if (diffHour < 24) {
    return `${diffHour} giờ trước`;
  } else if (diffDay <= 3) {
    return `${diffDay} ngày trước`;
  } else {
    return date.toLocaleString("vi-VN", {
      hour: "2-digit",
      minute: "2-digit",
      day: "2-digit",
      month: "2-digit",
      year: "numeric",
    });
  }
}

function formatDateOnCard(dateStr) {
  const date = new Date(dateStr);
  return date.toLocaleString("vi-VN", {
    day: "2-digit",
    month: "2-digit",
  });
}
watch(tempStartDate, (newStart) => {
  if (!newStart) return;
  if (!tempEndDate.value) return;

  const start = new Date(newStart);
  const end = new Date(tempEndDate.value);

  if (start >= end) {
    const newEnd = new Date(start);
    newEnd.setDate(newEnd.getDate() + 1);
    tempEndDate.value = newEnd.toISOString().slice(0, 10);
  }
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
function tooltipDate(card) {
  if (!card) return "Thẻ chưa hết hạn";
  if (card.completed) return "Thẻ này đã hoàn tất";
  if (!card.end_date) return "Thẻ chưa hết hạn";

  const now = new Date();
  const end = new Date(card.end_date);

  if (now > end) return "Thẻ đã hết hạn";

  const diff = Math.floor((end.getTime() - now.getTime()) / (1000 * 60 * 60));
  if (diff <= 24) return `Thẻ sẽ hết hạn trong ${diff} giờ`;

  return "Thẻ chưa hết hạn";
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
  const list = lists.value.find((l) => l.id === selectedCard.value.list_id);
  const card = list?.cards.find((c) => c.ID === selectedCard.value.ID);
  if (card) {
    card.members = [...members.value];
  }
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
  const list = lists.value.find((l) => l.id === selectedCard.value.list_id);
  const card = list?.cards.find((c) => c.ID === selectedCard.value.ID);
  if (card) {
    card.members = [...members.value];
  }
}
const props = defineProps(["card"]);

const filter = reactive({
  keyword: "",
  completed: false,
  incompleted: false,
  noDueDate: false,
  dueTomorrow: false,
  overdue: false,
  expireNextWeek: false,
  expireNextMonth: false,
  noMembers: false,
  assignedToMe: false,
  memberIds: null,
});

const listmembers = ref([]);

onMounted(async () => {
  const { data } = await useFetch(`http://localhost:3001/api/boards/${boardId}/members`);
  listmembers.value = data.value || [];
});

const listuseroption = computed(() =>
  listmembers.value.map(m => ({
    value: m.ID,
    name: m.name,
    avatar: m.avatar,
    email: m.email
  }))
);

// watchEffect(() => {
//   console.log("listuseroption", listuseroption.value);
// });
function checkVisible(card) {
  if (
    filter.keyword &&
    !card.title.toLowerCase().includes(filter.keyword.toLowerCase())
  ) {
    return false;
  }

  if (filter.completed && !card.completed) return false;
  if (filter.incompleted && card.completed) return false;

  const now = new Date();
  const end = card.end_date ? new Date(card.end_date) : null;

  const matchDueFilters = [];

  if (filter.noDueDate) {
    matchDueFilters.push(!end);
  }

  if (filter.dueTomorrow) {
    if (end) {
      const tomorrow = new Date(now);
      tomorrow.setDate(tomorrow.getDate() + 1);
      tomorrow.setHours(0, 0, 0, 0);
      const dayAfterTomorrow = new Date(tomorrow);
      dayAfterTomorrow.setDate(dayAfterTomorrow.getDate() + 1);

      matchDueFilters.push(end >= tomorrow && end < dayAfterTomorrow);
    }
  }

  if (filter.overdue) {
    matchDueFilters.push(end && now > end && !card.completed);
  }

  if (filter.expireNextWeek && end) {
    const nowDate = new Date(now);
    const dayOfWeek = nowDate.getDay();
    const daysUntilNextMonday = (8 - dayOfWeek) % 7 || 7;
    const nextWeekStart = new Date(nowDate);
    nextWeekStart.setDate(nowDate.getDate() + daysUntilNextMonday);
    nextWeekStart.setHours(0, 0, 0, 0);
    const nextWeekEnd = new Date(nextWeekStart);
    nextWeekEnd.setDate(nextWeekStart.getDate() + 6);
    nextWeekEnd.setHours(23, 59, 59, 999);
    matchDueFilters.push(end >= nextWeekStart && end <= nextWeekEnd);
  }

  if (filter.expireNextMonth && end) {
    const nextMonthStart = new Date(now.getFullYear(), now.getMonth() + 1, 1);
    const nextMonthEnd = new Date(now.getFullYear(), now.getMonth() + 2, 0);
    nextMonthEnd.setHours(23, 59, 59, 999);
    matchDueFilters.push(end >= nextMonthStart && end <= nextMonthEnd);
  }

  if (
    filter.noDueDate ||
    filter.dueTomorrow ||
    filter.expireNextWeek ||
    filter.expireNextMonth ||
    filter.overdue
  ) {
    if (!matchDueFilters.some(Boolean)) return false;
  }


if (!filter.noMembers && !filter.assignedToMe && !filter.memberIds) {
  return true;
}

const currentId = Number(currentUserId.value);
let matchMember = false;

if (filter.noMembers) {
  matchMember ||= !card.members || card.members.length === 0;
}

if (filter.assignedToMe) {
  matchMember ||= card.members?.some((m) => Number(m.UserID) === currentId) ?? false;
}

if (filter.memberIds?.length) {
  matchMember ||= card.members?.some((m) =>
    filter.memberIds.includes(m.user?.id ?? m.UserID)
  ) ?? false;
}


return matchMember;

}

watch(
  () => filter.completed,
  (val) => {
    if (val) filter.incompleted = false;
  }
);

watch(
  () => filter.incompleted,
  (val) => {
    if (val) filter.completed = false;
  }
);
watch(
  () => filter.dueTomorrow,
  (val) => {
    if (val) {
      (filter.expireNextWeek = false), (filter.expireNextMonth = false);
    }
  }
);
watch(
  () => filter.expireNextWeek,
  (val) => {
    if (val) {
      (filter.dueTomorrow = false), (filter.expireNextMonth = false);
    }
  }
);
watch(
  () => filter.expireNextMonth,
  (val) => {
    if (val) {
      (filter.dueTomorrow = false), (filter.expireNextWeek = false);
    }
  }
);
</script>
