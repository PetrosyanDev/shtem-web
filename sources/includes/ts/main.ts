import 'primeicons/primeicons.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import PrimeVue from 'primevue/config';
import ToastService from 'primevue/toastservice';

import MainQuiz from '@/views/MainQuiz.vue'
import GlobalQuiz from '@/views/GlobalQuiz.vue'

const quiz = createApp(MainQuiz)
quiz.use(createPinia())
quiz.mount('#question')

const quizBuilder = createApp(GlobalQuiz)
quizBuilder.use(createPinia())
quizBuilder.use(PrimeVue);
quizBuilder.use(ToastService);
quizBuilder.mount('#quiz')