<template>
  <div class="dashboard-wrapper">
    <div v-if="!currentUser" class="auth-container card">
      <header class="main-header">
        <h1>Wilber's Habit Tracker <span>{{ isRegistering ? 'Join' : 'Login' }}</span></h1>
      </header>
      
      <div class="auth-form">
        <input v-model="authEmail" placeholder="Email Address" class="auth-input" />
        <input v-model="authPassword" type="password" placeholder="Password" class="auth-input" />
        <button @click="handleAuth" class="btn-primary auth-btn">
          {{ isRegistering ? 'Create Account' : 'Sign In' }}
        </button>
        <p @click="isRegistering = !isRegistering" class="toggle-link">
          {{ isRegistering ? 'Already have an account? Login' : 'Need an account? Sign Up' }}
        </p>
      </div>
    </div>

    <div v-else class="container">
      <header class="main-header header-flex">
        <h1>Wilber's <span>Dashboard</span></h1>
        <button @click="logout" class="btn-logout">Logout</button>
      </header>

      <div class="progress-container card">
        <div class="progress-header">
          <span class="label">Daily Progress</span>
          <span class="value">{{ Math.round(progressPercentage) }}%</span>
        </div>
        <div class="progress-track">
          <div class="progress-fill" :style="{ width: progressPercentage + '%' }"></div>
        </div>
      </div>

      <div class="input-section card">
        <select v-model="selectedIcon" class="icon-select">
          <option v-for="emoji in emojiPalette" :key="emoji" :value="emoji">{{ emoji }}</option>
        </select>
        <input v-model="newHabit" @keyup.enter="addHabit" placeholder="What's your goal today?" />
        <button @click="addHabit" class="btn-primary">Add</button>
      </div>

      <div class="habit-list">
        <div v-for="habit in habits" :key="habit.id" class="habit-item card" :class="{ 'is-completed': habit.completed }">
          <div class="habit-content">
            <div class="habit-icon">{{ habit.icon }}</div>
            <div class="habit-text">
              <strong>{{ habit.name }}</strong>
            </div>
          </div>
          <div class="habit-actions">
            <button class="btn-done" @click="toggleHabit(habit)">{{ habit.completed ? 'Undo' : 'Done' }}</button>
            <button class="btn-delete" @click="deleteHabit(habit.id)">🗑️</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch, computed } from 'vue'
import axios from 'axios'

// --- Auth State ---
const currentUser = ref(JSON.parse(localStorage.getItem('user')))
const isRegistering = ref(false)
const authEmail = ref('')
const authPassword = ref('')

// --- Habit State ---
const habits = ref([])
const newHabit = ref('')
const selectedIcon = ref('🌱')
const emojiPalette = ["🌱", "💧", "🏃", "📚", "🧘", "💻", "🚗", "🪥", "🚿", "🙏", "🍳", "🛌"]

// --- Auth Logic ---
const handleAuth = async () => {
  const endpoint = isRegistering.value ? 'register' : 'login'
  try {
    const res = await axios.post(`http://localhost:3000/api/${endpoint}`, {
      email: authEmail.value,
      password: authPassword.value
    })
    currentUser.value = res.data
    localStorage.setItem('user', JSON.stringify(res.data))
    fetchHabits()
  } catch (err) {
    alert("Authentication failed. Check your credentials.")
  }
}

const logout = () => {
  currentUser.value = null
  localStorage.removeItem('user')
  habits.value = []
}

// --- Habit Logic ---
const fetchHabits = async () => {
  if (!currentUser.value) return
  const res = await axios.get(`http://localhost:3000/api/habits?userId=${currentUser.value.id}`)
  habits.value = res.data
}

const addHabit = async () => {
  if (!newHabit.value) return
  await axios.post('http://localhost:3000/api/habits', {
    name: newHabit.value,
    icon: selectedIcon.value,
    user_id: currentUser.value.id
  })
  newHabit.value = ''
  fetchHabits()
}

const toggleHabit = async (habit) => {
  await axios.put(`http://localhost:3000/api/habits/${habit.id}`)
  fetchHabits()
}

const deleteHabit = async (id) => {
  await axios.delete(`http://localhost:3000/api/habits/${id}`)
  fetchHabits()
}

const progressPercentage = computed(() => {
  if (habits.value.length === 0) return 0
  return (habits.value.filter(h => h.completed).length / habits.value.length) * 100
})

onMounted(fetchHabits)
</script>

<style scoped>
.dashboard-wrapper { min-height: 100vh; background-color: #f3f4f6; padding: 40px 20px; font-family: 'Inter', sans-serif; }
.container { max-width: 600px; margin: 0 auto; }
.main-header h1 { color: #333; font-weight: 700; }
.main-header span { color: #10b981; font-weight: 300; }
.card { background: white; border-radius: 12px; padding: 20px; box-shadow: 0 4px 6px -1px rgba(0,0,0,0.1); margin-bottom: 20px; }
.input-section { display: flex; gap: 10px; align-items: center; }
.icon-select {
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 8px;
  background: white;
  font-size: 1.2rem;
  cursor: pointer;
  outline: none;
}

.icon-select:focus {
  border-color: #10b981;
}
input { flex: 1; padding: 12px; border: 1px solid #ddd; border-radius: 8px; font-size: 1rem; }
.stats-grid { display: grid; grid-template-columns: 1fr 1fr; gap: 20px; }
.stat-card { text-align: center; }
.value { font-size: 1.8rem; font-weight: bold; display: block; }
.active { color: #10b981; }
.habit-item { display: flex; justify-content: space-between; align-items: center; border-left: 6px solid #10b981; transition: all 0.2s ease; }
.is-completed { border-left-color: #9ca3af; opacity: 0.6; transform: scale(0.98); }
.habit-content { display: flex; align-items: center; gap: 15px; }
.habit-icon { font-size: 2rem; }
.habit-text p { font-size: 0.85rem; color: #666; margin: 0; }
.habit-actions { display: flex; gap: 10px; }
.btn-primary, .btn-done { background: #10b981; color: white; border: none; padding: 10px 20px; border-radius: 8px; font-weight: 600; cursor: pointer; }
.btn-delete { background: #fee2e2; border: none; padding: 8px; border-radius: 8px; cursor: pointer; transition: 0.2s; }
.btn-delete:hover { background: #fecaca; }
.btn-reset {
  background-color: #f59e0b; /* Amber/Orange color */
  color: white;
  border: none;
  padding: 8px 12px;
  border-radius: 6px;
  cursor: pointer;
  margin-top: 5px;
  font-weight: bold;
}
.btn-reset:hover { background-color: #d97706; }
.progress-container {
  margin-top: 10px;
  padding: 15px 20px;
}

.progress-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}

.progress-track {
  width: 100%;
  height: 12px;
  background-color: #e5e7eb;
  border-radius: 10px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background: linear-gradient(90deg, #34d399, #10b981);
  transition: width 0.5s ease-out;
  box-shadow: 0 0 8px rgba(16, 185, 129, 0.3);
}
</style>