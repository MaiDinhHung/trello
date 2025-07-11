<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50">
    <UCard class="p-6 max-w-md w-full">
    <h1 class="text-2xl mb-4">Đăng nhập</h1>
    <UInput v-model="email" placeholder="Email" class="input mb-2 w-full" /><br>
    <UInput v-model="password" type="password" placeholder="Password" class="input mb-4 w-full" /><br>
    <UButton class="btn w-full" @click="login">Đăng nhập</UButton>
    </UCard>
  </div>
</template>

<script setup>
const email = ref("")
const password = ref("")
const router = useRouter()

async function login() {
    if (!email.value || !password.value) {
    alert("Vui lòng nhập email và mật khẩu")
    return
}
const { data, error } = await useFetch('http://localhost:3001/api/auth/login', {
    method: 'POST',
    body: { email: email.value, password: password.value }
})
    if (error.value) return alert("Sai tài khoản")

localStorage.setItem("token", data.value.token)
localStorage.setItem("user", JSON.stringify(data.value.user))
router.push('/boards/1')
}
</script>
