<template>
    <div v-if="bajins != null" class="container-fluid global-container">
        <!-- Header -->
        <h5>Ընտրեք բաժին</h5>
        <!-- Radio button group -->
        <div class="ms-4">
            <div v-for="bajin in bajins" :key="bajin.number">
                <input
                    class="form-check-input me-1"
                    type="radio"
                    :id="'flexRadioDefault' + bajin.number"
                    :value="bajin.number"
                    v-model="selectedBajin"
                />
                <label class="form-check-label" :for="'flexRadioDefault' + bajin.number">{{ bajin.name }}</label>
            </div>
            <label v-if="!isBajinSelected && !selectedBajin" class="mt-2 form-label text-danger">Բաժինը ընտրված չէ</label>
        </div>

        <!-- Advanced accordion -->
        <div class="accordion mt-3" id="advancedAccordion">
            <div class="accordion-item">
                <h2 class="accordion-header" id="advancedHeader">
                    <button
                        class="accordion-button"
                        type="button"
                        data-bs-toggle="collapse"
                        data-bs-target="#advancedCollapse"
                        aria-expanded="true"
                        aria-controls="advancedCollapse"
                    >
                        Advanced
                    </button>
                </h2>
                <div
                    id="advancedCollapse"
                    class="accordion-collapse collapse show"
                    aria-labelledby="advancedHeader"
                    data-bs-parent="#advancedAccordion"
                >
                    <div class="accordion-body">
                        <div class="form-check form-switch mb-2">
                            <input class="form-check-input" type="checkbox" id="randomSwitch" v-model="randomSwitch" />
                            <label class="form-check-label" for="randomSwitch">Պատահական</label>
                        </div>
                        <div class="form-check form-switch mb-2">
                            <input class="form-check-input" type="checkbox" id="skippableSwitch" v-model="skippableSwitch" />
                            <label class="form-check-label" for="skippableSwitch">Հնարավոր է բաց թողնել հարցը</label>
                        </div>
                        <div class="form-check form-switch mb-2">
                            <input class="form-check-input" type="checkbox" id="showNumberSwitch" v-model="showNumberSwitch" />
                            <label class="form-check-label" for="showNumberSwitch">Ցույց տալ թեստի համարը</label>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <div class="col col-md-6 col-lg-4 mx-auto">
            <button type="button" class="mt-3 btn btn-primary text-white w-100" style="border-radius: 50px" :onclick="startQuiz">
                Սկսել
            </button>
        </div>
    </div>
    <div v-else class="container-fluid global-container">
        <div class="w-100 d-flex justify-content-center">
            <div class="spinner-border text-primary" role="status">
                <span class="sr-only">Loading...</span>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

const isBajinSelected = ref(true)

const selectedBajin = ref()
const randomSwitch = ref(false)
const skippableSwitch = ref(false)
const showNumberSwitch = ref(true)
const bajins = ref()

const shtemName: string = window.location.href.split('/')[window.location.href.split('/').length - 2] || ''

function loadBajins() {
    const opts = {
        method: 'POST',
        headers: {
            'X-Shtem-Api-Key': 'someKey'
        }
    }

    fetch('https://shtemaran.am/api/v1/shtems/get-shtem-bajin/' + shtemName, opts)
        .then((responce) => responce.json())
        .then((data) => {
            if (data.error) {
                return
            }
            // console.log(data.data)
            bajins.value = data.data
        })
        .catch((error) => {
            console.log(error)
        })
}

const startQuiz = () => {
    // TODO CHECK IF CHOSEN

    if (!selectedBajin.value) {
        isBajinSelected.value = false
        return
    }

    // Define the type for query parameters
    interface QueryParams {
        bajin: string
        random: string
        skippable: string
        sn: string
    }

    // Extracting values for better readability
    const shtemUrl = 'https://shtemaran.am/shtems/'
    // const shtemUrl = 'http://localhost:9999/shtems/'
    const quizPath = '/quiz/'
    const queryParams: QueryParams = {
        bajin: selectedBajin.value,
        random: String(randomSwitch.value),
        skippable: String(skippableSwitch.value),
        sn: String(showNumberSwitch.value)
    }

    // Constructing the query string
    const queryString = new URLSearchParams(queryParams as unknown as Record<string, string>).toString()

    // Combining all parts to form the new URL
    const newUrl = `${shtemUrl}${shtemName}${quizPath}?${queryString}`

    // Redirecting to the new URL
    window.location.href = newUrl
}

onMounted(() => {
    loadBajins()
})
</script>

<style scoped>
/* Add scoped styles here if needed */
</style>
