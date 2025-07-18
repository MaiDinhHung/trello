<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50">
    <UCard class="p-6 max-w-md w-full">
      <h1 class="text-2xl font-bold pb-5 text-center">Đăng ký</h1>
      <label for="">Tên người dùng:</label>
      <UInput v-model="name" placeholder="Tên" class="mb-3 w-full" /><br>
      <label for="">Email:</label>
      <UInput v-model="email" placeholder="Email" class="mb-3 w-full" /><br>
      <label for="">Mật khẩu:</label>
      <UInput v-model="password" type="password" placeholder="Mật khẩu" class="mb-6 w-full" /><br>
      <UButton block @click="register" class="mb-6">Đăng ký</UButton>
      <ULink href="http://localhost:3000/login">Bạn đã có tài khoản? Đăng nhập ngay.</ULink>
    </UCard>
  </div>
</template>

<script setup>
const name = ref("")
const email = ref("")
const password = ref("")
const router = useRouter()
const toast = useToast()

async function register() {
  if (!name.value || !email.value || !password.value) {
    toast.add({ title: 'Vui lòng nhập đầy đủ thông tin', color: 'red' })
    return
  }
  const { error } = await useFetch('http://localhost:3001/api/auth/register', {
    method: 'POST',
    body: { name: name.value, email: email.value, password: password.value }
  })
  if (error.value) return toast.add({ title: 'Tài khoản đã tồn tại!', color: 'red' })
toast.add({ title: 'Đăng kí thành công', color: 'primary' })
  router.push('/login')
}
</script>
