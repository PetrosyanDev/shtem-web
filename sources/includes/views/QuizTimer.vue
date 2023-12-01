<template>
    <div class="d-flex gap-1 align-items-center">
        <img style="width: 20px; height: 20px" :src="'/assets/svg/stopwatch.svg'" alt="" />
        <p class="m-0">{{ formattedTime }}</p>
    </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'

const counter = ref({ min: 0, sec: 0 })

const formattedTime = computed(() => {
    const minutes = String(counter.value.min).padStart(2, '0')
    const seconds = String(counter.value.sec).padStart(2, '0')
    return `${minutes}:${seconds}`
})

const startStopwatch = () => {
    setInterval(() => {
        counter.value.sec += 1

        if (counter.value.sec === 60) {
            counter.value.sec = 0
            counter.value.min += 1
        }
    }, 1000)
}

onMounted(() => {
    startStopwatch()
})
</script>

<style></style>
