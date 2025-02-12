<template>
    <Loading v-if="isCallLoading"></Loading>
    <div v-else-if="callExists">
        <MeetingSetup @setup-complete="isSetupComplete = true" v-if="!isSetupComplete" :call="call" />
        <MeetingRoom :call="call" v-else />
    </div>

    <div class="h-screen flex items-center justify-center" v-else>
        <p class="text-2xl font-semibold">Meeting not found</p>
    </div>
</template>

<script setup>
import Loading from '@/components/Loading.vue'
import MeetingRoom from '@/components/MeetingRoom.vue'
import MeetingSetup from '@/components/MeetingSetup.vue'
import { useStreamStore } from '@/stores/stream'
import { ref } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const callID = route.params.id

let call = {}
const callExists = ref(false)
const isCallLoading = ref(true)
const isSetupComplete = ref(false)

const getCall = async (callID) => {
    try {
        const client = useStreamStore().client
        const { calls } = await client.queryCalls({ filter_conditions: { id: callID } })
        if (calls.length > 0) {
            call = calls[0]
            callExists.value = true
        }
    } catch (err) {
        console.log(err)
    } finally {
        isCallLoading.value = false
    }
}

getCall(callID)
</script>
