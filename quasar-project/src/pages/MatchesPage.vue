<template>
  <q-page padding>
    <div class="row q-col-gutter-md">
      <!-- Header -->
      <div class="col-12">
        <q-card>
          <q-card-section>
            <div class="row items-center justify-between">
              <div class="row items-center q-col-gutter-md">
                <div class="col-auto">
                  <q-icon name="sports_score" size="lg" class="text-primary" />
                </div>
                <div class="col">
                  <div class="text-h4">Matches</div>
                  <div class="text-subtitle1">Create and manage dart matches</div>
                </div>
              </div>
              <div>
                <q-btn 
                  color="primary" 
                  icon="add" 
                  label="New Match"
                  @click="showCreateMatchDialog = true"
                />
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>

      <!-- Match Filters -->
      <div class="col-12">
        <q-card>
          <q-card-section>
            <div class="row q-col-gutter-md items-center">
              <div class="col-12 col-md-3">
                <q-select
                  v-model="statusFilter"
                  :options="statusOptions"
                  label="Filter by Status"
                  clearable
                />
              </div>
              <div class="col-12 col-md-3">
                <q-select
                  v-model="playerFilter"
                  :options="playerOptions"
                  option-label="name"
                  option-value="id"
                  label="Filter by Player"
                  clearable
                  use-input
                  @filter="filterPlayers"
                />
              </div>
              <div class="col-12 col-md-6">
                <q-input
                  v-model="searchQuery"
                  label="Search matches"
                  clearable
                  debounce="300"
                />
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>

      <!-- Matches List -->
      <div class="col-12">
        <q-card>
          <q-card-section>
            <div class="text-h6 q-mb-md">Recent Matches</div>
            
            <q-list v-if="filteredMatches.length" separator>
              <q-item 
                v-for="match in filteredMatches" 
                :key="match.id" 
                clickable 
                @click="goToMatch(match.id)"
                class="q-pa-md"
              >
                <q-item-section avatar>
                  <q-avatar 
                    :color="getStatusColor(match.status)" 
                    text-color="white" 
                    :icon="getStatusIcon(match.status)"
                  />
                </q-item-section>
                
                <q-item-section>
                  <q-item-label class="text-h6">
                    {{ getPlayerName(match.player1_id) }} vs {{ getPlayerName(match.player2_id) }}
                  </q-item-label>
                  <q-item-label caption>
                    Status: {{ match.status }} | 
                    Score: {{ match.player1_score }} - {{ match.player2_score }}
                    <span v-if="match.tournament_id">| Tournament Match</span>
                  </q-item-label>
                  <q-item-label caption v-if="match.started_at">
                    Started: {{ formatDate(match.started_at) }}
                  </q-item-label>
                </q-item-section>

                <q-item-section side>
                  <div class="q-gutter-sm">
                    <q-chip 
                      :color="getStatusColor(match.status)" 
                      text-color="white" 
                      size="sm"
                    >
                      {{ match.status }}
                    </q-chip>
                    <q-btn 
                      v-if="match.status === 'pending'"
                      color="primary" 
                      size="sm" 
                      label="Start"
                      @click.stop="quickStartMatch(match)"
                    />
                    <q-btn 
                      v-if="match.status === 'in_progress'"
                      color="secondary" 
                      size="sm" 
                      label="Continue"
                      @click.stop="goToMatch(match.id)"
                    />
                  </div>
                </q-item-section>
              </q-item>
            </q-list>

            <div v-else-if="loading" class="text-center q-pa-lg">
              <q-spinner size="lg" color="primary" />
              <div class="text-subtitle1 q-mt-md">Loading matches...</div>
            </div>

            <div v-else class="text-center q-pa-lg">
              <q-icon name="sports_score" size="4rem" class="text-grey-5" />
              <div class="text-h6 text-grey-5 q-mt-md">No matches found</div>
              <div class="text-body2 text-grey-5">Create your first match to get started</div>
            </div>
          </q-card-section>
        </q-card>
      </div>
    </div>

    <!-- Create Match Dialog -->
    <q-dialog v-model="showCreateMatchDialog" persistent>
      <q-card style="min-width: 400px">
        <q-card-section>
          <div class="text-h6">Create New Match</div>
        </q-card-section>

        <q-card-section>
          <div class="q-gutter-md">
            <q-select
              v-model="newMatch.player1"
              :options="availablePlayers"
              option-label="name"
              option-value="id"
              label="Player 1"
              use-input
              @filter="filterPlayersForMatch"
              clearable
            />

            <q-select
              v-model="newMatch.player2"
              :options="availablePlayers"
              option-label="name"
              option-value="id"
              label="Player 2"
              use-input
              @filter="filterPlayersForMatch"
              clearable
            />

            <q-select
              v-model="newMatch.tournament"
              :options="availableTournaments"
              option-label="name"
              option-value="id"
              label="Tournament (Optional)"
              clearable
            />

            <q-checkbox
              v-model="newMatch.startImmediately"
              label="Start match immediately"
            />
          </div>
        </q-card-section>

        <q-card-actions align="right">
          <q-btn 
            flat 
            label="Cancel" 
            @click="cancelCreateMatch"
          />
          <q-btn 
            flat 
            label="Create" 
            color="primary" 
            @click="createMatch"
            :disable="!newMatch.player1 || !newMatch.player2 || newMatch.player1.id === newMatch.player2.id"
            :loading="loading"
          />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </q-page>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useQuasar } from 'quasar'
import { useMatchesStore } from 'src/stores/matches'
import { useTournamentsStore } from 'src/stores/tournaments'
import { playerApi } from 'src/services/api'

const router = useRouter()
const $q = useQuasar()
const matchesStore = useMatchesStore()
const tournamentsStore = useTournamentsStore()

// Reactive data
const showCreateMatchDialog = ref(false)
const players = ref([])
const availablePlayers = ref([])
const availableTournaments = ref([])
const statusFilter = ref(null)
const playerFilter = ref(null)
const searchQuery = ref('')

const newMatch = ref({
  player1: null,
  player2: null,
  tournament: null,
  startImmediately: false
})

// Computed
const loading = computed(() => matchesStore.loading)
const matches = computed(() => matchesStore.matches)

const statusOptions = ['pending', 'in_progress', 'completed']

const playerOptions = computed(() => players.value)

const filteredMatches = computed(() => {
  let filtered = matches.value

  // Filter by status
  if (statusFilter.value) {
    filtered = filtered.filter(match => match.status === statusFilter.value)
  }

  // Filter by player
  if (playerFilter.value) {
    filtered = filtered.filter(match => 
      match.player1_id === playerFilter.value.id || 
      match.player2_id === playerFilter.value.id
    )
  }

  // Search filter
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(match => {
      const player1Name = getPlayerName(match.player1_id).toLowerCase()
      const player2Name = getPlayerName(match.player2_id).toLowerCase()
      return player1Name.includes(query) || player2Name.includes(query)
    })
  }

  return filtered.sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
})

// Lifecycle
onMounted(async () => {
  await Promise.all([
    fetchPlayers(),
    fetchTournaments(),
    fetchMatches()
  ])
})

// Methods
async function fetchPlayers() {
  try {
    const response = await playerApi.getPlayers()
    players.value = Array.isArray(response.data.data) ? response.data.data : []
    availablePlayers.value = [...players.value]
  } catch (error) {
    console.error('Error fetching players:', error)
  }
}

async function fetchTournaments() {
  try {
    // You might want to implement a getTournaments API call
    availableTournaments.value = []
  } catch (error) {
    console.error('Error fetching tournaments:', error)
  }
}

async function fetchMatches() {
  // Since we don't have a general getMatches endpoint, 
  // we'll start with an empty list and populate as matches are created
  matchesStore.setMatches([])
}

function filterPlayers(val, update) {
  update(() => {
    if (val === '') {
      playerOptions.value = players.value
    } else {
      const needle = val.toLowerCase()
      playerOptions.value = players.value.filter(player => 
        player.name.toLowerCase().includes(needle)
      )
    }
  })
}

function filterPlayersForMatch(val, update) {
  update(() => {
    if (val === '') {
      availablePlayers.value = players.value
    } else {
      const needle = val.toLowerCase()
      availablePlayers.value = players.value.filter(player => 
        player.name.toLowerCase().includes(needle)
      )
    }
  })
}

function getPlayerName(playerId) {
  const player = players.value.find(p => p.id === playerId)
  return player?.name || 'Unknown Player'
}

function getStatusColor(status) {
  switch (status) {
    case 'pending': return 'orange'
    case 'in_progress': return 'blue'
    case 'completed': return 'green'
    default: return 'grey'
  }
}

function getStatusIcon(status) {
  switch (status) {
    case 'pending': return 'schedule'
    case 'in_progress': return 'play_arrow'
    case 'completed': return 'check_circle'
    default: return 'help'
  }
}

function formatDate(dateString) {
  if (!dateString) return 'N/A'
  return new Date(dateString).toLocaleString()
}

async function createMatch() {
  try {
    const matchData = {
      player1_id: newMatch.value.player1.id,
      player2_id: newMatch.value.player2.id
    }

    if (newMatch.value.tournament) {
      matchData.tournament_id = newMatch.value.tournament.id
    }

    const createdMatch = await matchesStore.createMatch(matchData)

    $q.notify({
      type: 'positive',
      message: 'Match created successfully'
    })

    showCreateMatchDialog.value = false

    // Start immediately if requested
    if (newMatch.value.startImmediately) {
      await quickStartMatch(createdMatch)
    }

  } catch (error) {
    $q.notify({
      type: 'negative',
      message: 'Failed to create match'
    })
  }
}

async function quickStartMatch(match) {
  try {
    await matchesStore.startMatch(match.id)
    
    $q.notify({
      type: 'positive',
      message: 'Match started successfully'
    })

    // Navigate to match detail
    router.push(`/matches/${match.id}`)
  } catch (error) {
    $q.notify({
      type: 'negative',
      message: 'Failed to start match'
    })
  }
}

function goToMatch(matchId) {
  router.push(`/matches/${matchId}`)
}

function cancelCreateMatch() {
  showCreateMatchDialog.value = false
  newMatch.value = {
    player1: null,
    player2: null,
    tournament: null,
    startImmediately: false
  }
}
</script>