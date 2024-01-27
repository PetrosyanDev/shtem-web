<template>
    <div class="d-flex gap-1 align-items-center">
        <img style="width: 20px; height: 20px" :src="'/assets/svg/stopwatch.svg'" alt="" />
        <p class="m-0">{{ time }}</p>
    </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue'

export default defineComponent({
    props: {
        hours: Boolean,
        minutes: Boolean
    },
    setup(props, { emit }) {
        const time = ref<string | null>(null)
        const isRunning = ref(false)
        const startTime = ref<number | null>(null)
        const times = ref<number[]>([0, 0, 0, 0])
        let frameId: number | null = null

        const start = () => {
            if (isRunning.value) throw new Error('Stopwatch has already started.')
            isRunning.value = true
            if (!startTime.value) startTime.value = performance.now()
            frameId = requestAnimationFrame(step)
        }

        const lap = (id: number) => {
            emit('lap', performance.now(), time.value, id)
        }

        const stop = () => {
            if (!isRunning.value) throw new Error('Stopwatch has not been started yet.')
            isRunning.value = false
            startTime.value = null
            times.value = [0, 0, 0, 0]
            cancelAnimationFrame(frameId!)
        }

        const reset = () => {
            startTime.value = 0
            isRunning.value = false
            times.value = [0, 0, 0, 0]
            time.value = formatTimes()
        }

        const formatTimes = (): string => {
            const pad0 = (value: number, count: number): string => {
                let result = value.toString()
                while (result.length < count) {
                    result = '0' + result
                    --count
                }
                return result
            }

            const hours = pad0(times.value[0], 2)
            const minutes = pad0(times.value[1], 2)
            const seconds = pad0(times.value[2], 2)
            const milliseconds = pad0(Math.trunc(times.value[3] % 100), 2)

            if (props.hours) {
                return `${hours}:${minutes}:${seconds}:${milliseconds}`
            }

            if (props.minutes) {
                return `${minutes}:${seconds}`
            }

            return `${seconds}:${milliseconds}`
        }

        const step = (timestamp: number) => {
            if (!isRunning.value) return
            calculate(timestamp)
            startTime.value = timestamp
            time.value = formatTimes()
            frameId = requestAnimationFrame(step)
        }

        const calculate = (timestamp: number) => {
            const diff = timestamp - (startTime.value as number)
            times.value[3] += diff / 10
            if (times.value[3] >= 100) {
                times.value[3] -= 100
                times.value[2] += 1
            }
            if (times.value[2] >= 60) {
                times.value[2] -= 60
                times.value[1] += 1
            }
            if (times.value[1] >= 60) {
                times.value[1] -= 60
                times.value[0] += 1
            }
        }

        onMounted(() => {
            start()
        })

        return {
            time,
            isRunning,
            start,
            lap,
            stop,
            reset
        }
    }
})
</script>
