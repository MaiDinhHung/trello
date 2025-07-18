<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50">
    <UCard class="p-8 max-w-md w-full space-y-6 rounded-xl shadow-md">
      <div class="flex items-center justify-center gap-3 mb-6">
        <!-- <UButton size="xl" color="secondary" icon="i-lucide-folder-open" variant="ghost" /> -->
        <h1 class="font-bold text-4xl text-gray-800">Trello</h1>
      </div>

      <div class="flex flex-col gap-3 mb-3">
        <UButton
        class="mb-2"
          label="Đăng nhập"
          icon="i-lucide-log-in"
          color="primary"
          to="http://localhost:3000/login"
          target="_blank"
        />

        <UButton
          label="Đăng ký"
          color="neutral"
          variant="outline"
          icon="i-lucide-user-plus"
          to="http://localhost:3000/register"
          target="_blank"
        />
      </div>
    </UCard>  
  </div>
</template>

