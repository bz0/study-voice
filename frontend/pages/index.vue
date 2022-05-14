<template>
  <div class="wrap">
    <p
      v-if="username"
      class="message"
    >
      ようこそ {{ username }} さん
    </p>
    <VueRecordAudio @result="onResult" />
    <button
      v-if="loggedIn"
      class="button"
      @click="logout"
    >
      ログアウト
    </button>
    <button
      v-else
      class="button"
      @click="login"
    >
      ログイン
    </button>
  </div>
</template>

<script>
import Vue from 'vue'
import VueRecord from '@codekraft-studio/vue-record'
Vue.use(VueRecord)

export default {
  name: 'IndexPage',
  computed: {
    username () {
      if (this.$auth.$state.user &&
          this.$auth.$state.user.name) {
        console.log('state', this.$auth.$state)
        return this.$auth.$state.user.name
      } else {
        return null
      }
    },
    loggedIn () {
      console.log(this.$auth)
      return this.$auth.loggedIn
    }
  },
  methods: {
    login () {
      this.$auth.loginWith('auth0')
    },
    logout () {
      this.$auth.logout()
    },
    async onResult (data) {
      console.log('The blob data:', data)
      console.log('Downloadable audio', window.URL.createObjectURL(data))

      const uploadBlobData = new FormData()
      uploadBlobData.append('data', new Blob([data], { type: 'video/webm' }))

      const url = 'http://localhost:8080/record'
      const response = await this.$axios.$post(url, uploadBlobData, {
        headers: {
          'content-type': 'multipart/form-data'
        }
      })

      console.log('response', response)
    }
  }
}
</script>
