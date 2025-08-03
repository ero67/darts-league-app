<template>
  <q-page padding>
    <div class="row q-col-gutter-md">
      <div class="col-12">
        <q-card class="bg-primary text-white">
          <q-card-section>
            <div class="row items-center q-col-gutter-md">
              <div class="col-auto">
                <q-icon name="sports" size="3rem" />
              </div>
              <div class="col">
                <div class="text-h3">Darts League</div>
                <div class="text-subtitle1">Welcome to your darts league management system</div>
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>

      <div class="col-12 col-md-6 col-lg-3">
        <q-card>
          <q-card-section>
            <div class="row items-center q-col-gutter-md">
              <div class="col-auto">
                <q-icon name="people" size="2rem" class="text-blue" />
              </div>
              <div class="col">
                <div class="text-h5">{{ playersStore.players.length }}</div>
                <div class="text-subtitle2">Players</div>
              </div>
            </div>
          </q-card-section>
          <q-card-actions>
            <q-btn flat color="primary" label="Manage" @click="$router.push('/players')" />
          </q-card-actions>
        </q-card>
      </div>

      <div class="col-12 col-md-6 col-lg-3">
        <q-card>
          <q-card-section>
            <div class="row items-center q-col-gutter-md">
              <div class="col-auto">
                <q-icon name="emoji_events" size="2rem" class="text-orange" />
              </div>
              <div class="col">
                <div class="text-h5">{{ leaguesStore.leagues.length }}</div>
                <div class="text-subtitle2">Leagues</div>
              </div>
            </div>
          </q-card-section>
          <q-card-actions>
            <q-btn flat color="primary" label="Manage" @click="$router.push('/leagues')" />
          </q-card-actions>
        </q-card>
      </div>

      <div class="col-12 col-md-6 col-lg-3">
        <q-card>
          <q-card-section>
            <div class="row items-center q-col-gutter-md">
              <div class="col-auto">
                <q-icon name="tournament" size="2rem" class="text-purple" />
              </div>
              <div class="col">
                <div class="text-h5">{{ activeTournaments }}</div>
                <div class="text-subtitle2">Active Tournaments</div>
              </div>
            </div>
          </q-card-section>
          <q-card-actions>
            <q-btn flat color="primary" label="View" @click="$router.push('/tournaments')" />
          </q-card-actions>
        </q-card>
      </div>

      <div class="col-12 col-md-6 col-lg-3">
        <q-card>
          <q-card-section>
            <div class="row items-center q-col-gutter-md">
              <div class="col-auto">
                <q-icon name="sports_score" size="2rem" class="text-green" />
              </div>
              <div class="col">
                <div class="text-h5">{{ activeMatches }}</div>
                <div class="text-subtitle2">Active Matches</div>
              </div>
            </div>
          </q-card-section>
          <q-card-actions>
            <q-btn flat color="primary" label="View" @click="$router.push('/matches')" />
          </q-card-actions>
        </q-card>
      </div>

      <div class="col-12 col-lg-8">
        <q-card>
          <q-card-section>
            <div class="text-h6">Recent Activity</div>
          </q-card-section>
          <q-card-section class="q-pt-none">
            <q-list>
              <q-item v-if="recentActivity.length === 0">
                <q-item-section>
                  <q-item-label class="text-center text-grey-5">No recent activity</q-item-label>
                </q-item-section>
              </q-item>
              <q-item v-for="activity in recentActivity" :key="activity.id">
                <q-item-section avatar>
                  <q-icon :name="activity.icon" :color="activity.color" />
                </q-item-section>
                <q-item-section>
                  <q-item-label>{{ activity.title }}</q-item-label>
                  <q-item-label caption>{{ activity.description }}</q-item-label>
                </q-item-section>
                <q-item-section side>
                  <q-item-label caption>{{ activity.time }}</q-item-label>
                </q-item-section>
              </q-item>
            </q-list>
          </q-card-section>
        </q-card>
      </div>

      <div class="col-12 col-lg-4">
        <q-card>
          <q-card-section>
            <div class="text-h6">Quick Actions</div>
          </q-card-section>
          <q-card-section class="q-pt-none">
            <q-list>
              <q-item clickable @click="$router.push('/players')">
                <q-item-section avatar>
                  <q-icon name="person_add" color="primary" />
                </q-item-section>
                <q-item-section>
                  <q-item-label>Add New Player</q-item-label>
                </q-item-section>
              </q-item>
              <q-item clickable @click="$router.push('/leagues')">
                <q-item-section avatar>
                  <q-icon name="add_circle" color="primary" />
                </q-item-section>
                <q-item-section>
                  <q-item-label>Create League</q-item-label>
                </q-item-section>
              </q-item>
              <q-item clickable @click="$router.push('/tournaments')">
                <q-item-section avatar>
                  <q-icon name="tournament" color="primary" />
                </q-item-section>
                <q-item-section>
                  <q-item-label>Start Tournament</q-item-label>
                </q-item-section>
              </q-item>
              <q-item clickable @click="$router.push('/standings')">
                <q-item-section avatar>
                  <q-icon name="leaderboard" color="primary" />
                </q-item-section>
                <q-item-section>
                  <q-item-label>View Standings</q-item-label>
                </q-item-section>
              </q-item>
            </q-list>
          </q-card-section>
        </q-card>
      </div>
    </div>
  </q-page>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { usePlayersStore } from 'src/stores/players'
import { useLeaguesStore } from 'src/stores/leagues'
import { useTournamentsStore } from 'src/stores/tournaments'
import { useMatchesStore } from 'src/stores/matches'

const playersStore = usePlayersStore()
const leaguesStore = useLeaguesStore()
const tournamentsStore = useTournamentsStore()
const matchesStore = useMatchesStore()

const recentActivity = ref([])

const activeTournaments = computed(() => {
  return tournamentsStore.activeTournaments.length
})

const activeMatches = computed(() => {
  return matchesStore.activeMatches.length
})

onMounted(() => {
  playersStore.fetchPlayers()
  leaguesStore.fetchLeagues()
  
  // Mock recent activity for demo
  recentActivity.value = [
    {
      id: 1,
      icon: 'person_add',
      color: 'blue',
      title: 'New player registered',
      description: 'John Doe joined the league',
      time: '2 hours ago'
    },
    {
      id: 2,
      icon: 'emoji_events',
      color: 'orange',
      title: 'League started',
      description: 'Spring League 2024 is now active',
      time: '1 day ago'
    },
    {
      id: 3,
      icon: 'sports_score',
      color: 'green',
      title: 'Match completed',
      description: 'Alice vs Bob - Alice won 3-1',
      time: '2 days ago'
    }
  ]
})
</script>
