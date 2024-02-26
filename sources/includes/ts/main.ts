import { createApp } from 'vue'
import { createPinia } from 'pinia'

import MainQuiz from '@/views/MainQuiz.vue'
import GlobalQuiz from '@/views/GlobalQuiz.vue'

const quiz = createApp(MainQuiz)
quiz.use(createPinia())
quiz.mount('#question')

const quizBuilder = createApp(GlobalQuiz)
quizBuilder.use(createPinia())
quizBuilder.mount('#quiz')