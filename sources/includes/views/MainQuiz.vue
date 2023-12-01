<template>
    <div class="container-fluid">
        <!-- <div class="d-flex justify-content-center">Some</div> -->
        <div class="d-flex justify-content-center"><QuizTimer></QuizTimer></div>
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
                    <p>Picked ->{{ q_answer }}</p>
                    <div class="mt-4 question-answer row gap-2">
                        <div v-for="(item, index) in answers" :key="item">
                            <input type="radio" :id="(index + 1).toString()" :value="index + 1" v-model="q_answer" />
                            <label :for="(index + 1).toString" class="col-12 btn border rounded">{{ item }} {{ index + 1 }}</label>
                        </div>
                    </div>

                    <button v-if="q_answer != ''" type="button" class="btn btn-primary">Next</button>
                </div>
            </div>
            <div class="d-none d-xl-flex col-1 flex-fill"></div>
        </div>
    </div>
</template>

<script setup lang="ts">
import QuizTimer from './QuizTimer.vue'

var q_answer = '1'
var num = 1
var text = 'Ո՞ր շարքի բոլոր բառերում է գրվում բ:'
var answers = ['answer', 'answer', 'answer', 'answer']
</script>

<style></style>
