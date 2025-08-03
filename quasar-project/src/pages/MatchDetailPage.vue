<template>
  <q-page padding>
    <div class="row q-col-gutter-md">
      <!-- Match Header -->
      <div class="col-12">
        <q-card>
          <q-card-section>
            <div class="row items-center justify-between">
              <div>
                <div class="text-h4" v-if="currentMatch">
                  {{ getPlayerName(currentMatch.player1_id) }} vs {{ getPlayerName(currentMatch.player2_id) }}
                </div>
                <div class="text-h4" v-else>Loading Match...</div>
                <div class="text-subtitle1 text-grey-6" v-if="currentMatch">
                  Status: {{ currentMatch.status }} | Round: {{ currentMatch.round }}
                </div>
              </div>
              <div class="q-gutter-sm">
                <q-btn 
                  v-if="currentMatch?.status === 'pending'"
                  color="primary" 
                  icon="play_arrow" 
                  label="Start Match"
                  @click="startMatch"
                  :loading="loading"
                />
                <q-btn 
                  v-if="currentMatch?.status === 'in_progress'"
                  color="negative" 
                  icon="stop" 
                  label="End Match"
                  @click="showCompleteMatchDialog = true"
                  :loading="loading"
                />
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>

      <!-- Game Setup (if match not started) -->
      <div class="col-12" v-if="currentMatch?.status === 'pending'">
        <q-card>
          <q-card-section>
            <div class="text-h6 q-mb-md">Game Setup</div>
            <div class="row q-col-gutter-md">
              <div class="col-12 col-md-4">
                <q-select
                  v-model="gameSettings.gameType"
                  :options="gameTypes"
                  label="Game Type"
                  option-label="label"
                  option-value="value"
                />
              </div>
              <div class="col-12 col-md-4">
                <q-input
                  v-model.number="gameSettings.startingScore"
                  type="number"
                  label="Starting Score"
                  :disable="gameSettings.gameType?.value !== '01'"
                />
              </div>
              <div class="col-12 col-md-4">
                <q-select
                  v-model="gameSettings.finishType"
                  :options="finishTypes"
                  label="Finish Type"
                  :disable="gameSettings.gameType?.value !== '01'"
                />
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>

      <!-- Dart Scoring Interface (if match in progress) -->
      <div class="col-12" v-if="currentMatch?.status === 'in_progress'">
        <q-card>
          <q-card-section>
            <!-- Current Player Turn -->
            <div class="text-center q-mb-md">
              <q-chip 
                :color="currentPlayerTurn === 'player1' ? 'primary' : 'grey'" 
                text-color="white" 
                size="lg"
                class="q-mr-md"
              >
                {{ getPlayerName(currentMatch.player1_id) }}'s Turn
              </q-chip>
              <q-chip 
                :color="currentPlayerTurn === 'player2' ? 'primary' : 'grey'" 
                text-color="white" 
                size="lg"
              >
                {{ getPlayerName(currentMatch.player2_id) }}'s Turn
              </q-chip>
            </div>

            <!-- Current Scores Display -->
            <div class="row q-col-gutter-md text-center q-mb-lg">
              <div class="col-6">
                <q-card flat bordered :class="currentPlayerTurn === 'player1' ? 'bg-primary text-white' : ''">
                  <q-card-section>
                    <div class="text-h3">{{ player1Score }}</div>
                    <div class="text-subtitle1">{{ getPlayerName(currentMatch.player1_id) }}</div>
                    <div class="text-caption">Sets: {{ player1Sets }} | Legs: {{ player1Legs }}</div>
                  </q-card-section>
                </q-card>
              </div>
              <div class="col-6">
                <q-card flat bordered :class="currentPlayerTurn === 'player2' ? 'bg-primary text-white' : ''">
                  <q-card-section>
                    <div class="text-h3">{{ player2Score }}</div>
                    <div class="text-subtitle1">{{ getPlayerName(currentMatch.player2_id) }}</div>
                    <div class="text-caption">Sets: {{ player2Sets }} | Legs: {{ player2Legs }}</div>
                  </q-card-section>
                </q-card>
              </div>
            </div>

            <!-- Current Throw Display -->
            <div class="text-center q-mb-md">
              <div class="text-h6">Current Throw</div>
              <div class="q-gutter-sm q-mb-md">
                <q-chip 
                  v-for="(dart, index) in currentThrow" 
                  :key="index" 
                  color="secondary" 
                  text-color="white"
                  size="lg"
                >
                  {{ dart || '-' }}
                </q-chip>
                <q-chip 
                  v-for="i in (3 - currentThrow.length)" 
                  :key="'empty-' + i" 
                  color="grey-3" 
                  text-color="grey-7"
                  size="lg"
                >
                  -
                </q-chip>
              </div>
              <div class="text-subtitle1">
                Total: {{ currentThrowTotal }} | Average: {{ currentAverage.toFixed(2) }}
              </div>
            </div>

            <!-- Dart Input -->
            <div class="row q-col-gutter-md q-mb-md">
              <div class="col-12 col-md-8 offset-md-2">
                <div class="text-center">
                  <div class="text-h6 q-mb-md">Enter Dart Score</div>
                  
                  <!-- Quick Score Buttons -->
                  <div class="row q-col-gutter-sm q-mb-md">
                    <div class="col-2" v-for="score in quickScores" :key="score">
                      <q-btn 
                        :label="score" 
                        color="primary" 
                        @click="addDartScore(score)"
                        :disable="currentThrow.length >= 3"
                        class="full-width"
                      />
                    </div>
                  </div>

                  <!-- Manual Score Input -->
                  <div class="row q-col-gutter-sm q-mb-md">
                    <div class="col-8">
                      <q-input
                        v-model="manualScore"
                        type="number"
                        label="Manual Score (0-180)"
                        min="0"
                        max="180"
                        @keyup.enter="addDartScore(manualScore)"
                      />
                    </div>
                    <div class="col-4">
                      <q-btn 
                        label="Add" 
                        color="secondary" 
                        @click="addDartScore(manualScore)"
                        :disable="!manualScore || currentThrow.length >= 3"
                        class="full-width"
                      />
                    </div>
                  </div>

                  <!-- Control Buttons -->
                  <div class="q-gutter-sm">
                    <q-btn 
                      label="Undo Last Dart" 
                      color="warning" 
                      @click="undoLastDart"
                      :disable="currentThrow.length === 0"
                    />
                    <q-btn 
                      label="Submit Throw" 
                      color="positive" 
                      @click="submitThrow"
                      :disable="currentThrow.length === 0"
                    />
                    <q-btn 
                      label="Bust" 
                      color="negative" 
                      @click="bustThrow"
                    />
                  </div>
                </div>
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>

      <!-- Throw History -->
      <div class="col-12" v-if="currentMatch?.status === 'in_progress'">
        <q-card>
          <q-card-section>
            <div class="text-h6 q-mb-md">Throw History</div>
            <div class="row">
              <div class="col-6">
                <div class="text-subtitle1 q-mb-sm">{{ getPlayerName(currentMatch.player1_id) }}</div>
                <q-list separator>
                  <q-item v-for="(throwData, index) in player1History" :key="index">
                    <q-item-section>
                      <q-item-label>{{ throwData.darts.join(', ') }}</q-item-label>
                      <q-item-label caption>Total: {{ throwData.total }} | Score: {{ throwData.score }}</q-item-label>
                    </q-item-section>
                  </q-item>
                </q-list>
              </div>
              <div class="col-6">
                <div class="text-subtitle1 q-mb-sm">{{ getPlayerName(currentMatch.player2_id) }}</div>
                <q-list separator>
                  <q-item v-for="(throwData, index) in player2History" :key="index">
                    <q-item-section>
                      <q-item-label>{{ throwData.darts.join(', ') }}</q-item-label>
                      <q-item-label caption>Total: {{ throwData.total }} | Score: {{ throwData.score }}</q-item-label>
                    </q-item-section>
                  </q-item>
                </q-list>
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>
    </div>

    <!-- Complete Match Dialog -->
    <q-dialog v-model="showCompleteMatchDialog">
      <q-card style="min-width: 300px">
        <q-card-section>
          <div class="text-h6">Complete Match</div>
        </q-card-section>
        <q-card-section>
          <div class="text-body1 q-mb-md">Select the winner:</div>
          <q-option-group
            v-model="selectedWinner"
            :options="winnerOptions"
            color="primary"
          />
        </q-card-section>
        <q-card-actions align="right">
          <q-btn flat label="Cancel" @click="showCompleteMatchDialog = false" />
          <q-btn 
            flat 
            label="Complete" 
            color="primary" 
            @click="completeMatch"
            :disable="!selectedWinner"
          />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </q-page>
</template>

<script setup>
import { ref, onMounted, computed, watch } from 'vue'
import { useRoute } from 'vue-router'
import { useQuasar } from 'quasar'
import { useMatchesStore } from 'src/stores/matches'
import { playerApi } from 'src/services/api'

const route = useRoute()
const $q = useQuasar()
const matchesStore = useMatchesStore()

const players = ref({})
const currentPlayerTurn = ref('player1')
const currentThrow = ref([])
const manualScore = ref('')
const showCompleteMatchDialog = ref(false)
const selectedWinner = ref(null)

// Game state
const player1Score = ref(501)
const player2Score = ref(501)
const player1Sets = ref(0)
const player2Sets = ref(0)
const player1Legs = ref(0)
const player2Legs = ref(0)
const player1History = ref([])
const player2History = ref([])
const totalDartsThrown = ref(0)

// Game settings
const gameSettings = ref({
  gameType: { label: '501', value: '01' },
  startingScore: 501,
  finishType: 'Double Out'
})

const gameTypes = [
  { label: '501', value: '01' },
  { label: '301', value: '01' },
  { label: 'Cricket', value: 'cricket' }
]

const finishTypes = ['Straight Out', 'Double Out', 'Master Out']

const quickScores = [0, 26, 45, 60, 85, 100, 140, 180]

const loading = computed(() => matchesStore.loading)
const currentMatch = computed(() => matchesStore.currentMatch)

const currentThrowTotal = computed(() => {
  return currentThrow.value.reduce((sum, dart) => sum + (parseInt(dart) || 0), 0)
})

const currentAverage = computed(() => {
  if (totalDartsThrown.value === 0) return 0
  const totalScore = (gameSettings.value.startingScore * 2) - player1Score.value - player2Score.value
  return totalScore / totalDartsThrown.value * 3
})

const winnerOptions = computed(() => {
  if (!currentMatch.value) return []
  return [
    {
      label: getPlayerName(currentMatch.value.player1_id),
      value: currentMatch.value.player1_id
    },
    {
      label: getPlayerName(currentMatch.value.player2_id),
      value: currentMatch.value.player2_id
    }
  ]
})

onMounted(async () => {
  const matchId = route.params.id
  await fetchMatchData(matchId)
  await fetchPlayers()
})

watch(() => gameSettings.value.startingScore, (newScore) => {
  if (currentMatch.value?.status === 'pending') {
    player1Score.value = newScore
    player2Score.value = newScore
  }
})

async function fetchMatchData(matchId) {
  try {
    await matchesStore.fetchMatch(matchId)
  } catch (error) {
    $q.notify({
      type: 'negative',
      message: 'Failed to load match data'
    })
  }
}

async function fetchPlayers() {
  try {
    const response = await playerApi.getPlayers()
    const playersList = Array.isArray(response.data.data) ? response.data.data : []
    players.value = {}
    playersList.forEach(player => {
      players.value[player.id] = player
    })
  } catch (error) {
    console.error('Error fetching players:', error)
  }
}

function getPlayerName(playerId) {
  return players.value[playerId]?.name || 'Unknown Player'
}

async function startMatch() {
  try {
    await matchesStore.startMatch(route.params.id)
    
    // Initialize game state
    player1Score.value = gameSettings.value.startingScore
    player2Score.value = gameSettings.value.startingScore
    player1Sets.value = 0
    player2Sets.value = 0
    player1Legs.value = 0
    player2Legs.value = 0
    player1History.value = []
    player2History.value = []
    totalDartsThrown.value = 0
    currentPlayerTurn.value = 'player1'
    
    $q.notify({
      type: 'positive',
      message: 'Match started successfully'
    })
  } catch (error) {
    $q.notify({
      type: 'negative',
      message: 'Failed to start match'
    })
  }
}

function addDartScore(score) {
  if (currentThrow.value.length >= 3 || !score) return
  
  const dartScore = parseInt(score)
  if (dartScore < 0 || dartScore > 180) return
  
  currentThrow.value.push(dartScore)
  manualScore.value = ''
}

function undoLastDart() {
  if (currentThrow.value.length > 0) {
    currentThrow.value.pop()
  }
}

function submitThrow() {
  if (currentThrow.value.length === 0) return
  
  const throwTotal = currentThrowTotal.value
  const currentPlayer = currentPlayerTurn.value
  const currentScore = currentPlayer === 'player1' ? player1Score.value : player2Score.value
  const newScore = currentScore - throwTotal
  
  // Check for bust (score goes below 0 or ends on 1 with double out)
  const isBust = newScore < 0 || 
    (gameSettings.value.finishType === 'Double Out' && newScore === 1) ||
    (newScore === 0 && gameSettings.value.finishType === 'Double Out' && !isDouble(currentThrow.value[currentThrow.value.length - 1]))
  
  if (isBust) {
    bustThrow()
    return
  }
  
  // Update score
  if (currentPlayer === 'player1') {
    player1Score.value = newScore
    player1History.value.push({
      darts: [...currentThrow.value],
      total: throwTotal,
      score: newScore
    })
  } else {
    player2Score.value = newScore
    player2History.value.push({
      darts: [...currentThrow.value],
      total: throwTotal,
      score: newScore
    })
  }
  
  totalDartsThrown.value += currentThrow.value.length
  
  // Check for leg win
  if (newScore === 0) {
    if (currentPlayer === 'player1') {
      player1Legs.value++
    } else {
      player2Legs.value++
    }
    
    $q.notify({
      type: 'positive',
      message: `${getPlayerName(currentMatch.value[currentPlayer + '_id'])} wins the leg!`
    })
    
    // Reset for new leg
    player1Score.value = gameSettings.value.startingScore
    player2Score.value = gameSettings.value.startingScore
  }
  
  // Switch player turn
  currentPlayerTurn.value = currentPlayerTurn.value === 'player1' ? 'player2' : 'player1'
  
  // Clear current throw
  currentThrow.value = []
}

function bustThrow() {
  const currentPlayer = currentPlayerTurn.value
  
  // Add bust to history
  if (currentPlayer === 'player1') {
    player1History.value.push({
      darts: [...currentThrow.value],
      total: currentThrowTotal.value,
      score: player1Score.value,
      bust: true
    })
  } else {
    player2History.value.push({
      darts: [...currentThrow.value],
      total: currentThrowTotal.value,
      score: player2Score.value,
      bust: true
    })
  }
  
  totalDartsThrown.value += currentThrow.value.length
  
  $q.notify({
    type: 'negative',
    message: 'BUST!'
  })
  
  // Switch player turn
  currentPlayerTurn.value = currentPlayerTurn.value === 'player1' ? 'player2' : 'player1'
  
  // Clear current throw
  currentThrow.value = []
}

function isDouble(score) {
  const doubles = [2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 50]
  return doubles.includes(parseInt(score))
}

async function completeMatch() {
  if (!selectedWinner.value) return
  
  try {
    await matchesStore.completeMatch(route.params.id, selectedWinner.value)
    
    $q.notify({
      type: 'positive',
      message: 'Match completed successfully'
    })
    
    showCompleteMatchDialog.value = false
  } catch (error) {
    $q.notify({
      type: 'negative',
      message: 'Failed to complete match'
    })
  }
}
</script>