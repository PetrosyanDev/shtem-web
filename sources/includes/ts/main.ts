import { createApp } from 'vue'
import { createPinia } from 'pinia'

import MainQuiz from '@/views/MainQuiz.vue'

const quiz = createApp(MainQuiz)
quiz.use(createPinia())
quiz.mount('#question')