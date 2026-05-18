<script setup>
import { reactive } from 'vue'

const state = reactive({
  mode: "login", // login | register

  // login
  email: "",
  password: "",

  // register
  name: "",
  regEmail: "",
  regPassword: "",

  message: ""
})

function login() {
  if (!state.email || !state.password) {
    state.message = "Preenche email e password"
    return
  }

  state.message = `Login tentado com ${state.email}`
}

function register() {
  if (!state.name || !state.regEmail || !state.regPassword) {
    state.message = "Preenche todos campos"
    return
  }

  state.message = `Conta criada para ${state.name}`
}

function toggleMode() {
  state.mode = state.mode === "login" ? "register" : "login"
  state.message = ""
}
</script>

<template>
  <main class="container">

    <h2 v-if="state.mode === 'login'">Login</h2>
    <h2 v-else>Cadastro</h2>

    <!-- LOGIN -->
    <div v-if="state.mode === 'login'">
      <input v-model="state.email" placeholder="Email" class="input" />
      <input v-model="state.password" type="password" placeholder="Password" class="input" />

      <button @click="login" class="btn">Entrar</button>
    </div>

    <!-- REGISTER -->
    <div v-else>
      <input v-model="state.name" placeholder="Nome" class="input" />
      <input v-model="state.regEmail" placeholder="Email" class="input" />
      <input v-model="state.regPassword" type="password" placeholder="Password" class="input" />

      <button @click="register" class="btn">Criar conta</button>
    </div>

    <!-- TOGGLE -->
    <button @click="toggleMode" class="link">
      {{ state.mode === 'login' ? 'Criar conta' : 'Já tenho conta' }}
    </button>

    <p class="result">{{ state.message }}</p>

  </main>
</template>

<style scoped>
.container {
  padding: 20px;
  max-width: 300px;
}

.input {
  display: block;
  padding: 10px;
  margin-bottom: 10px;
  width: 100%;
}

.btn {
  padding: 10px;
  width: 100%;
  cursor: pointer;
}

.link {
  margin-top: 10px;
  background: none;
  border: none;
  color: blue;
  cursor: pointer;
}

.result {
  margin-top: 15px;
}
</style>