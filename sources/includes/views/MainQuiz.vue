<template>
    <div class="container-fluid">
        <div class="row flex-nowrap">
            <div class="d-none d-xl-flex col-1 flex-fill"></div>
            <div class="col flex-fill global-container">
                <div class="col col-md-10 mx-auto bg-light bg-gradient rounded question-body">
                    <div class="w-100">
                        <div class="queston-title">
                            <QuizTimer class="position-timer"></QuizTimer>
                            <h4 class="text-center m-0">Բաժին {{ bajin }} Մաս {{ mas }} Համար {{ q_number }}</h4>
                        </div>
                        <div class="question-text">
                            <p>{{ text }}</p>
                        </div>
                    </div>
                    <div class="mt-4 question-answer row gap-2">
                        <div v-for="(item, index) in answers" :key="item">
                            <label v-if="q_answer === (index + 1).toString()" :for="(index + 1).toString()" class="col-12 btn border rounded border-success text-start">
                                <input type="radio" name="question" :id="(index + 1).toString()" :value="(index + 1).toString()" v-model="q_answer" style="display: none" />
                                {{ item }}
                            </label>
                            <label v-else :for="(index + 1).toString()" class="col-12 btn border rounded text-start">
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
import { ref, onBeforeMount } from 'vue'
import QuizTimer from './QuizTimer.vue'

var shtemaran = window.location.href.split('/').pop()
let bajin: number = 1
let mas: number = 1
let q_number: number = 1

onBeforeMount(() => {
    GetNextQuestion()
})

let text = ''
let answers = [
    '1) Տղամարդը անբարիացակամ վերաբերմունք դրսևորեց իր նոր հարևանների նկատմամբ:',
    '2) Բանջար քաղող աղջիկների բույլը վետվետում էր լանջերն ի վեր:',
    '3) Նա ամբողջ օրն անցկացնում էր թեթևաբարո կանանց և զեխ գինարբուքների մեջ:',
    '4) Գեղջկուհու` արևից ու քամուց թրծված դեմքը խոսում էր նրա հոգսաշատ կյանքի մասին:'
]
let q_answer = ref('')

const nextQuestion = () => {
    // Handle the logic for moving to the next question
    console.log('Selected Answer:', q_answer.value)

    q_answer.value = ''
    q_number += 1
}

function GetNextQuestion() {
    var payload = {
        shtemaran: shtemaran,
        bajin: bajin,
        mas: mas,
        number: q_number
    }

    const options = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'X-Shtem-Api-Key': 'someKey'
        },
        body: JSON.stringify(payload)
    }

    fetch('http://localhost:9998/api/v1/questions/find', options)
        .then((responce) => responce.json())
        .then((data) => {
            console.log(data)
        })
}
</script>

<style></style>
