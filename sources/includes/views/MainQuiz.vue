<template>
    <div class="container-fluid global-container">
        <div class="row flex-nowrap">
            <div class="col flex-fill">
                <div v-if="currentQuestion.text" class="col mx-auto bg-light bg-gradient rounded question-body">
                    <div class="w-100">
                        <div
                            class="queston-title d-flex flex-column flex-sm-row justify-content-between align-items-center gap-1"
                        >
                            <Stopwatch class="position-timer" :minutes="true" ref="stopwatchRef" />
                            <h4 v-if="showNumberSwitch" class="text-center m-0">
                                Բաժին {{ currentQuestion.bajin }} Մաս {{ currentQuestion.mas }} Համար
                                {{ currentQuestion.number }}
                            </h4>
                            <button
                                v-if="skippableSwitch"
                                class="d-none d-sm-block btn btn-primary text-white me-2"
                                @click="loadQuestion"
                                style="border-radius: 50px"
                            >
                                Skip
                            </button>
                        </div>

                        <div class="question-text mt-4">
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
                    <button
                        v-if="skippableSwitch"
                        class="w-100 mt-4 d-block d-sm-none btn btn-primary text-white me-2"
                        @click="loadQuestion"
                        style="border-radius: 50px"
                    >
                        Skip
                    </button>
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
import { ref, onMounted, onBeforeMount } from 'vue'
import Question from './Models'
import Stopwatch from './StopWatch.vue'

// PARAMETERS
const queryString = window.location.search
const params = new URLSearchParams(queryString)

const selectedBajin = Number(params.get('bajin'))
const randomSwitch = ref(params.get('random') === 'true')
const skippableSwitch = ref(params.get('skippable') === 'true')
const showNumberSwitch = ref(params.get('sn') === 'true')

// GLOBAL
const shtemName: string = window.location.href.split('/').slice(-2)[0] || ''
const currentQuestion = ref(new Question(shtemName, selectedBajin, 1, 0, '', [''], [0]))

const stopwatchRef = ref<InstanceType<typeof Stopwatch> | null>(null)
let canClick = true
const questionCounter = ref(0)

// QUESTIONS
let Questions: Question[] = []

function loadBajin() {
    canClick = false
    const payload = { shtemaran: currentQuestion.value.shtemaran, bajin: currentQuestion.value.bajin }
    const opts = {
        method: 'POST',
        headers: { 'X-Shtem-Api-Key': 'someKey' },
        body: JSON.stringify(payload)
    }

    fetch('https://shtemaran.am/api/v1/questions/findBajin', opts)
        .then((response) => response.json())
        .then((data) => {
            if (data.error) {
                currentQuestion.value.text = data.error
                currentQuestion.value.options = []
                currentQuestion.value.number += 1
            } else {
                Questions = data.data

                if (randomSwitch.value) {
                    shuffleQuestions()
                }

                loadQuestion()
            }
        })
        .catch((error) => {
            currentQuestion.value.text = error
            currentQuestion.value.options = []
            currentQuestion.value.number += 1
        })
}

const loadQuestion = () => {
    canClick = true
    itemsRef = []

    if (Questions.length > questionCounter.value) {
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
        showNumberSwitch.value = false
        skippableSwitch.value = false
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
        clearSelected(divContainer)
    } else {
        console.log('Cannot select option')
    }
}

onBeforeMount(() => {
    if (selectedBajin == 0) {
        window.location.href = `https://shtemaran.am/shtems/${shtemName}/build-quiz`
    }
})

onMounted(() => {
    loadBajin()
})

function shuffleQuestions() {
    for (let i = Questions.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1))
        ;[Questions[i], Questions[j]] = [Questions[j], Questions[i]]
    }
}
</script>

<style></style>
