<template>
    <div class="container-fluid global-container">
        <div class="row flex-nowrap">
            <div class="col flex-fill">
                <div v-if="currentQuestion.text" class="col mx-auto bg-light bg-gradient rounded question-body">
                    <div class="w-100">
                        <div class="queston-title">
                            <Stopwatch class="position-timer" :minutes="true" ref="stopwatchRef" />
                            <h4 class="text-center m-0">
                                Բաժին {{ currentQuestion.bajin }} Մաս {{ currentQuestion.mas }} Համար {{ currentQuestion.number }}
                            </h4>
                        </div>
                        <div class="question-text">
                            <div class="ql-editor" v-html="currentQuestion.text"></div>
                        </div>
                    </div>
                    <div class="mt-4 question-answer row gap-2">
                        <div class="question-answer" v-for="(choice, item) in currentQuestion.options" :key="item">
                            <label class="col-12 btn rounded border" :ref="optionChosen" @click="onOptionClicked(choice, item)">
                                {{ choice }}
                            </label>
                        </div>
                    </div>
                </div>
                <div v-else class="col mx-auto bg-light bg-gradient rounded question-body justify-content-center">
                    <div class="w-100 d-flex justify-content-center">
                        <div class="spinner-border text-default" role="status"></div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import Question from './Models'
import Stopwatch from './StopWatch.vue'

let canClick = true
let questionCounter = ref(0)

const shtemName: string = window.location.href.split('/')[window.location.href.split('/').length - 2] || ''
// const bajinCounter: string = window.location.href.split('/')[window.location.href.split('/').length - 2] || ''
const currentQuestion = ref(new Question(shtemName, 1, 1, 0, '', [''], [0]))
let Questions: Question[] = []

const stopwatchRef = ref<InstanceType<typeof Stopwatch> | null>(null)

function loadBajin() {
    canClick = true

    // Other

    var payload = {
        shtemaran: currentQuestion.value.shtemaran,
        bajin: currentQuestion.value.bajin,
        mas: 1,
        number: 1
    }

    const opts = {
        method: 'POST',
        headers: {
            'X-Shtem-Api-Key': 'someKey'
        },
        body: JSON.stringify(payload)
    }

    fetch('https://shtemaran.am/api/v1/questions/findBajin', opts)
        .then((responce) => responce.json())
        .then((data) => {
            if (data.error) {
                currentQuestion.value.text = data.error
                currentQuestion.value.options = []
                currentQuestion.value.number += 1
                return
            }
            Questions = data.data
            loadQuestion()
        })
        .catch((error) => {
            currentQuestion.value.number += 1
            currentQuestion.value.text = error
            currentQuestion.value.options = []
        })
}

const loadQuestion = () => {
    canClick = true

    itemsRef = []

    if (Questions.length > questionCounter.value) {
        // load question
        currentQuestion.value = Questions[questionCounter.value]
        questionCounter.value++
        if (stopwatchRef.value) {
            stopwatchRef.value.reset()
            stopwatchRef.value.start()
        }
    } else {
        currentQuestion.value.text = 'Out of questions'
        currentQuestion.value.options = []
        currentQuestion.value.number += 1
        if (stopwatchRef.value) {
            stopwatchRef.value.reset()
        }
    }
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
    }, 500)
}

const onOptionClicked = (choice: any, item: any) => {
    if (canClick) {
        const divContainer = itemsRef[item]
        const optionID = item + 1
        if (currentQuestion.value.answers[0] == optionID) {
            divContainer.classList.add('border-success')
        } else {
            divContainer.classList.add('border-danger')
            setTimeout(() => {
                divContainer.classList.remove('border-danger')
            }, 500)
            return
        }
        canClick = false
        // TODO go to next question
        clearSelected(divContainer)
    } else {
        // Cant select option
        console.log('cant select question')
    }
}

onMounted(() => {
    loadBajin()
})
</script>

<style></style>
