<template>
    <div class="container-fluid">
        <div class="row flex-nowrap">
            <div class="d-none d-xl-flex col-1 flex-fill"></div>
            <div class="col flex-fill global-container">
                <div v-if="currentQuestion.text" class="col col-md-10 mx-auto bg-light bg-gradient rounded question-body">
                    <div class="w-100">
                        <div class="queston-title">
                            <QuizTimer class="position-timer"></QuizTimer>
                            <h4 class="text-center m-0">Բաժին {{ currentQuestion.bajin }} Մաս {{ currentQuestion.mas }} Համար {{ currentQuestion.q_number }}</h4>
                        </div>
                        <div class="question-text">
                            <p>{{ currentQuestion.text }}</p>
                        </div>
                    </div>
                    <div class="mt-4 question-answer row gap-2">
                        <div v-for="(choice, item) in currentQuestion.options" :key="item">
                            <label class="col-12 btn rounded border" :ref="optionChosen" @click="onOptionClicked(choice, item)"> {{ choice }}{{ item }} </label>
                        </div>
                    </div>
                </div>
                <div v-else class="col col-md-10 mx-auto bg-light bg-gradient rounded question-body justify-content-center">
                    <div class="w-100 d-flex justify-content-center">
                        <div class="spinner-border text-default" role="status"></div>
                    </div>
                </div>
            </div>
            <div class="d-none d-xl-flex col-1 flex-fill"></div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import QuizTimer from './QuizTimer.vue'

const currentQuestion = ref({
    shtemaran: window.location.href.split('/').pop(),
    bajin: 1,
    mas: 1,
    q_number: 0,
    text: '',
    options: [''],
    answers: [0]
})

let canClick = true

function loadQuestion() {
    canClick = true

    // Other

    var payload = {
        shtemaran: currentQuestion.value.shtemaran,
        bajin: currentQuestion.value.bajin,
        mas: currentQuestion.value.mas,
        number: currentQuestion.value.q_number + 1
    }

    const opts = {
        method: 'POST',
        headers: {
            'X-Shtem-Api-Key': 'someKey'
        },
        body: JSON.stringify(payload)
    }

    console.log(opts.body)

    fetch('https://shtemaran.am/api/v1/questions/find', opts)
        .then((responce) => responce.json())
        .then((data) => {
            if (data.error) {
                currentQuestion.value.text = data.error
                currentQuestion.value.options = []
                return
            }
            currentQuestion.value.text = data.data.text
            currentQuestion.value.options = data.data.options
            currentQuestion.value.answers = data.data.answers
            currentQuestion.value.q_number += 1
        })
        .catch((error) => {
            currentQuestion.value.text = error
            currentQuestion.value.options = []
        })
}

let itemsRef: any = []
const optionChosen = (element: any) => {
    if (element) {
        itemsRef.push(element)
    }
}

const clearSelected = (divSelected: any) => {
    setTimeout(() => {
        divSelected.classList.remove('border-success')
        divSelected.classList.remove('border-danger')
        loadQuestion()
    }, 1000)
}

const onOptionClicked = (choice: any, item: any) => {
    if (canClick) {
        const divContainer = itemsRef[item + 1]
        const optionID = item + 1
        console.log(currentQuestion.value.answers[0], optionID)
        if (currentQuestion.value.answers[0] == optionID) {
            console.log('you are correct')
            divContainer.classList.add('border-success')
        } else {
            console.log('you are wrong')
            divContainer.classList.add('border-danger')
            setTimeout(() => {
                divContainer.classList.remove('border-danger')
            }, 500)
            return
        }
        canClick = false
        // TODO go to next question
        clearSelected(divContainer)
        console.log(choice, item)
    } else {
        // Cant select option
        console.log('cant select question')
    }
}

onMounted(() => {
    loadQuestion()
})

// const nextQuestion = () => {
//     console.log('Selected Answer:', q_answer.value)

//     if (q_answer.value == answers[0].toString()) {
//         q_answer.value = ''
//         GetQuestion()
//     } else {
//         const element = document.querySelector('.border-success.selectedd')
//         if (element) {
//             element.classList.remove('border-success')
//             element.classList.add('border-danger')
//         }
//     }
// }
</script>

<style></style>
