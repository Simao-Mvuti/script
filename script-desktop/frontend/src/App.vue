<script setup>
import { reactive } from 'vue'
import { CadastrarUsuario } from "../../frontend/wailsjs/go/main/App"

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

  state.message = `Login feito ${state.email}`
}
async function register() {

  if (!state.name || !state.regEmail || !state.regPassword) {
    state.message = "Preenche todos os campos"
    return
  }

  try {

    const resultado = await CadastrarUsuario({
      nome: state.name,
      email: state.regEmail,
      senha: state.regPassword
    })

    console.log(resultado)

    state.message = "Conta criada com sucesso"

  } catch (err) {

    console.error(err)

    state.message = String(err)
  }
}

function toggleMode() {
  state.mode = state.mode === "login" ? "register" : "login"
  state.message = ""
}
</script>

<template>
  <div class="page">

    <div class="card">

      <h1 class="title">
        {{ state.mode === "login" ? "Bem vindo de volta" : "Criar onta" }}
      </h1>

      <p class="subtitle">
        {{ state.mode === "login" ? "Logar para continuar" : "Juntos Hoje" }}
      </p>

      <!-- LOGIN -->
      <div v-if="state.mode === 'login'">

        <input v-model="state.email" type="email" placeholder="Email" class="input" />
        <input v-model="state.password" type="password" placeholder="Password" class="input" />

        <button @click="login" class="button">Sign in</button>

      </div>

      <!-- REGISTER -->
      <div v-else>

        <input v-model="state.name" type="text" placeholder="Nome" class="input" />
        <input v-model="state.regEmail" type="email" placeholder="Email" class="input" />
        <input v-model="state.regPassword" type="password" placeholder="Password" class="input" />

        <button @click="register" class="button">Create account</button>

      </div>

      <!-- TOGGLE -->
      <button @click="toggleMode" class="link">
        {{ state.mode === "login" ? "Criar conta" : "Já tenho conta" }}
      </button>

      <p class="message">{{ state.message }}</p>

    </div>

  </div>
</template>

<style scoped>
.page {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: radial-gradient(circle at top, #1e3a8a, #0f172a);
  font-family: system-ui;
}

.card {
  width: 340px;
  padding: 30px;
  border-radius: 16px;

  background: rgba(255, 255, 255, 0.08);
  backdrop-filter: blur(12px);

  border: 1px solid rgba(255, 255, 255, 0.15);
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.4);

  text-align: center;
}

.title {
  color: white;
  font-size: 22px;
  margin-bottom: 5px;
}

.subtitle {
  color: rgba(255,255,255,0.6);
  font-size: 13px;
  margin-bottom: 20px;
}

.input {
  width: 100%;
  padding: 12px;
  margin-bottom: 10px;

  border: none;
  border-radius: 10px;

  background: rgba(255,255,255,0.1);
  color: white;

  outline: none;
}

.input::placeholder {
  color: rgba(255,255,255,0.5);
}

.button {
  width: 100%;
  padding: 12px;

  border: none;
  border-radius: 10px;

  background: #3b82f6;
  color: white;

  cursor: pointer;
  transition: 0.2s;
}

.button:hover {
  background: #2563eb;
  transform: translateY(-2px);
}

.link {
  margin-top: 10px;
  background: none;
  border: none;
  color: #60a5fa;
  cursor: pointer;
}

.message {
  margin-top: 12px;
  color: rgba(255,255,255,0.7);
  font-size: 13px;
}
</style>