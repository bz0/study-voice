<template>
  <div class="wrap">
    <p
      v-if="username"
      class="message"
    >
      ようこそ {{ username }} さん
    </p>
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
export default {
  name: 'IndexPage',
  computed: {
    username () {
      if (this.$auth.$state.user &&
          this.$auth.$state.user.name) {
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
    }
  }
}
</script>
