<template>
    <div class="container-fluid global-container">
        <!-- Header -->
        <h5>Ընտրեք բաժին</h5>

        <!-- Radio button group -->
        <div class="ms-4">
            <div v-for="bajin in bajins" :key="bajin.number">
                <input
                    class="form-check-input me-1"
                    type="radio"
                    :id="'flexRadioDefault' + bajin.number"
                    :value="bajin.name"
                    v-model="selectedBajin"
                />
                <label class="form-check-label" :for="'flexRadioDefault' + bajin.number">{{ bajin.name }}</label>
            </div>
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
            <button data-v-610dfcc7="" type="button" class="mt-3 btn btn-primary text-white w-100" style="border-radius: 50px">
                Սկսել
            </button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const selectedBajin = ref()
const randomSwitch = ref(false)
const skippableSwitch = ref(false)
const showNumberSwitch = ref(false)

const bajins = [
    { name: 'Հնչյունափոխություն', number: 1 },
    { name: 'Բառագիտություն', number: 2 },
    { name: 'Ձևաբանություն', number: 3 },
    { name: 'Շարահյուսություն', number: 4 },
    { name: 'Կապակցված խոսք', number: 5 }
]
</script>

<style scoped>
/* Add scoped styles here if needed */
</style>
