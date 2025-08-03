<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated class="bg-primary">
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer"
        />

        <q-toolbar-title>
          <div class="row items-center">
            <q-icon name="sports" size="sm" class="q-mr-sm" />
            Darts League
          </div>
        </q-toolbar-title>

        <q-btn
          flat
          dense
          round
          icon="refresh"
          aria-label="Refresh"
          @click="checkBackendStatus"
        />
        <q-chip
          :color="backendStatus === 'online' ? 'positive' : 'negative'"
          text-color="white"
          size="sm"
          :label="backendStatus"
        />
      </q-toolbar>
    </q-header>

    <q-drawer
      v-model="leftDrawerOpen"
      show-if-above
      bordered
      class="bg-grey-1"
    >
      <q-list>
        <q-item-label
          header
          class="text-grey-8"
        >
          Navigation
        </q-item-label>

        <q-item
          v-for="link in navigationLinks"
          :key="link.title"
          :to="link.route"
          clickable
          v-ripple
          class="q-mb-xs"
        >
          <q-item-section avatar>
            <q-icon :name="link.icon" />
          </q-item-section>

          <q-item-section>
            <q-item-label>{{ link.title }}</q-item-label>
            <q-item-label caption>{{ link.caption }}</q-item-label>
          </q-item-section>
        </q-item>

        <q-separator class="q-my-md" />

        <q-item-label
          header
          class="text-grey-8"
        >
          Quick Actions
        </q-item-label>

        <q-item
          clickable
          v-ripple
          @click="$router.push('/players/create')"
        >
          <q-item-section avatar>
            <q-icon name="person_add" />
          </q-item-section>
          <q-item-section>
            <q-item-label>Add Player</q-item-label>
          </q-item-section>
        </q-item>

        <q-item
          clickable
          v-ripple
          @click="$router.push('/leagues/create')"
        >
          <q-item-section avatar>
            <q-icon name="add_circle" />
          </q-item-section>
          <q-item-section>
            <q-item-label>New League</q-item-label>
          </q-item-section>
        </q-item>
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { healthApi } from 'src/services/api'

const navigationLinks = [
  {
    title: 'Dashboard',
    caption: 'Overview and stats',
    icon: 'dashboard',
    route: '/'
  },
  {
    title: 'Players',
    caption: 'Manage players',
    icon: 'people',
    route: '/players'
  },
  {
    title: 'Leagues',
    caption: 'Manage leagues',
    icon: 'emoji_events',
    route: '/leagues'
  },
  {
    title: 'Tournaments',
    caption: 'Tournament brackets',
    icon: 'tournament',
    route: '/tournaments'
  },
  {
    title: 'Matches',
    caption: 'Match results',
    icon: 'sports_score',
    route: '/matches'
  },
  {
    title: 'Standings',
    caption: 'League standings',
    icon: 'leaderboard',
    route: '/standings'
  }
]

const leftDrawerOpen = ref(false)
const backendStatus = ref('checking')

function toggleLeftDrawer () {
  leftDrawerOpen.value = !leftDrawerOpen.value
}

async function checkBackendStatus() {
  try {
    await healthApi.check()
    backendStatus.value = 'online'
  } catch (error) {
    backendStatus.value = 'offline'
  }
}

onMounted(() => {
  checkBackendStatus()
})
</script>
