<template>
    <div class="container-fluid">
        <!-- <div class="d-flex justify-content-center">Some</div> -->
        <div class="d-flex justify-content-center"><QuizTimer></QuizTimer></div>
        <div class="row flex-nowrap">
            <div class="d-none d-xl-flex col-1 flex-fill"></div>
            <div class="col flex-fill global-container">
                <div class="mt-5 w-100 bg-light bg-gradient rounded question-body">
                    <div class="question-title">
                        <h2>Համար {{ num }}</h2>
                    </div>
                    <div class="question-text">
                        <p>{{ text }}</p>
                    </div>
                </div>
            </div>
            <div class="d-none d-xl-flex col-1 flex-fill"></div>
        </div>
    </div>
</template>

<script setup lang="ts">
import QuizTimer from './QuizTimer.vue'

var num = 1
var text = 'Ո՞ր շարքի բոլոր բառերում է գրվում բ:'
</script>

<style></style>
