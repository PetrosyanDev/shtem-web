<template>
    <div class="container-fluid">
        <div class="d-flex justify-content-center">
            <QuizTimer></QuizTimer>
        </div>
        <div class="row flex-nowrap">
            <div class="d-none d-xl-flex col-1 flex-fill"></div>
            <div class="col flex-fill global-container">
                <div class="col col-md-10 mx-auto bg-light bg-gradient rounded question-body">
                    <div class="w-100">
                        <div class="question-title">
                            <h4>Համար {{ num }}</h4>
                        </div>
                        <div class="question-text">
                            <p>{{ text }}</p>
                        </div>
                    </div>
                    <div class="mt-4 question-answer row gap-2">
                        <div v-for="(item, index) in answers" :key="item">
                            <label v-if="q_answer === (index + 1).toString()" :for="(index + 1).toString()" class="col-12 btn border rounded border-success text-start" :style="{ 'padding-left': textsPadding }">
                                <input type="radio" name="question" :id="(index + 1).toString()" :value="(index + 1).toString()" v-model="q_answer" style="display: none" />
                                {{ item }}
                            </label>
                            <label v-else :for="(index + 1).toString()" class="col-12 btn border rounded text-start" :style="{ 'padding-left': textsPadding }">
                                <input type="radio" name="question" :id="(index + 1).toString()" :value="(index + 1).toString()" v-model="q_answer" style="display: none" />
                                {{ item }}
                            </label>
                        </div>
                        <div v-if="q_answer !== ''" class="d-flex justify-content-center">
                            <button @click="nextQuestion" type="button" class="col-12 col-md-6 col-lg-4 btn btn-primary mt-3">Next</button>
                        </div>
                    </div>
                </div>
            </div>
            <div class="d-none d-xl-flex col-1 flex-fill"></div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import QuizTimer from './QuizTimer.vue'

let q_answer = ref('')
let num = 1
let text = 'Ո՞ր նախադասության մեջ ուղղագրական սխալ ունեցող բառ կա:'
let answers = [
    '1) Տղամարդը անբարիացակամ վերաբերմունք դրսևորեց իր նոր հարևանների նկատմամբ:',
    '2) Բանջար քաղող աղջիկների բույլը վետվետում էր լանջերն ի վեր:',
    '3) Նա ամբողջ օրն անցկացնում էր թեթևաբարո կանանց և զեխ գինարբուքների մեջ:',
    '4) Գեղջկուհու` արևից ու քամուց թրծված դեմքը խոսում էր նրա հոգսաշատ կյանքի մասին:'
]

const nextQuestion = () => {
    // Handle the logic for moving to the next question
    console.log('Selected Answer:', q_answer.value)

    q_answer.value = ''
    num += 1
    calculatePadding()
}

var textsPadding = ref('0px')

const calculatePadding = () => {
    const maxLength = answers.reduce((max, answer) => {
        const length = answer.length
        return length > max ? length : max
    }, 0)

    if (maxLength > 110) {
        textsPadding.value = '0px'
    } else if (maxLength > 80) {
        textsPadding.value = '20px'
    } else if (maxLength > 50) {
        textsPadding.value = '20px'
    } else if (maxLength > 30) {
        textsPadding.value = '10px'
    } else {
        textsPadding.value = '0px'
    }
}

calculatePadding()
</script>

<style></style>
