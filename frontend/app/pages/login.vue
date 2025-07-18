<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50">
    <UCard class="p-6 max-w-md w-full space-y-4">
      <h1 class="text-2xl font-bold pb-5 text-center">Đăng nhập</h1>
      <label for="" class="text-start">Tài khoản:</label>
      <UInput v-model="email" placeholder="Email" class="w-full mb-3" /><br>
      <label for="">Mật khẩu:</label>
      <UInput v-model="password" type="password" placeholder="Mật khẩu" class="w-full mb-4" /><br>
      <UButton block @click="login" class="mb-6">Đăng nhập</UButton><br>  
      <ULink href="http://localhost:3000/register">Bạn chưa có tài khoản? Đăng kí ngay.</ULink>
    </UCard>  
  </div>
</template>

<script setup>
const email = ref("")
const password = ref("")
const router = useRouter()
const toast = useToast()

async function login() {
  if (!email.value || !password.value) {
    toast.add({ title: 'Vui lòng nhập email và mật khẩu', color: 'red' })
    return
  }

  const { data, error } = await useFetch('http://localhost:3001/api/auth/login', {
    method: 'POST',
    body: { email: email.value, password: password.value }
  })

  if (error.value || !data.value?.token) {
    toast.add({ title: 'Sai tài khoản hoặc mật khẩu', color: 'red' })
    return
  }

  localStorage.setItem('token', data.value.token)
  localStorage.setItem('user', JSON.stringify(data.value.user))

  toast.add({ title: 'Đăng nhập thành công', color: 'green' })
  router.push('/boards/1')
}
</script>
