<template>
  <div>
    <q-card>
      <q-card-section class="flex justify-between">
        <q-input class="col-grow" dense @keydown.enter="msgSend" outlined v-model="input" label="Сообщение"/>
        <q-btn outline color="primary" icon-right="send" @click="msgSend">SEND</q-btn>
      </q-card-section>
    </q-card>
    <q-card>
      <q-card-section>
        <q-chat-message v-for="(msg, i) in messages" :key="i" :text="[msg.text]"/>
      </q-card-section>
    </q-card>
  </div>
</template>

<script>
import {ref} from 'vue'
import API from './api'

export default {
  name: 'App',

  components: {},

  setup() {
    const messages = ref([])
    const input = ref('')
    const msgSend = ref(() => {
      API.sendMessage(input.value)
      input.value = ''
    })

    let startMessageCycle = () => {
    }
    startMessageCycle = async (id = 0) => {
      const msg = await API.getMessage(id)
      if (msg.length === 0) {
        setTimeout(async () => {
          await startMessageCycle(id)
        }, 500)
      } else {
        if (msg.text !== '') messages.value.push(msg)
        setTimeout(async () => {
          await startMessageCycle(id + 1)
        }, 100)
      }
    }
    startMessageCycle()

    return {
      messages,
      input,
      msgSend
    }
  }
}
</script>
