<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50">
    <UCard class="p-6 max-w-md w-full">
      <h1 class="text-2xl mb-4 text-center">Đăng ký</h1>
      <UInput v-model="name" placeholder="Tên" class="mb-2 w-full" /><br>
      <UInput v-model="email" placeholder="Email" class="mb-2 w-full" /><br>
      <UInput v-model="password" type="password" placeholder="Mật khẩu" class="mb-4 w-full" /><br>
      <UButton block @click="register">Đăng ký</UButton>
    </UCard>
  </div>
</template>


<script setup>
const name = ref("")
const email = ref("")
const password = ref("")
const router = useRouter()

async function register() {
  const { error } = await useFetch('http://localhost:3001/api/auth/register', {
    method: 'POST',
    body: { name: name.value, email: email.value, password: password.value }
  })
  if (error.value) return alert("Lỗi đăng ký")

  router.push('/login')
}
</script>
