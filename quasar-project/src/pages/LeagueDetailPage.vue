<template>
  <q-page padding>
    <div class="row q-col-gutter-md">
      <!-- League Info Header -->
      <div class="col-12">
        <q-card>
          <q-card-section>
            <div class="row items-center justify-between">
              <div>
                <div class="text-h4" v-if="currentLeague">
                  {{ currentLeague.name }}
                </div>
                <div class="text-h4" v-else>Loading...</div>
                <div class="text-subtitle1 text-grey-6" v-if="currentLeague">
                  Status: {{ currentLeague.status }} | Players:
                  {{ currentLeague.players?.length || 0 }}
                </div>
              </div>
              <div class="q-gutter-sm">
                <q-btn
                  v-if="currentLeague?.status === 'setup'"
                  color="primary"
                  icon="play_arrow"
                  label="Start League"
                  @click="startLeague"
                  :loading="loading"
                />
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>

      <!-- Add Player Section -->
      <div class="col-12 col-md-6">
        <q-card>
          <q-card-section>
            <div class="text-h6 q-mb-md">
              <q-icon name="person_add" class="q-mr-sm" />
              Add Player to League
            </div>

            <q-select
              v-model="selectedPlayer"
              :options="availablePlayers"
              option-label="name"
              option-value="id"
              label="Select Player"
              use-input
              input-debounce="300"
              @filter="filterPlayers"
              clearable
              class="q-mb-md"
            >
              <template v-slot:no-option>
                <q-item>
                  <q-item-section class="text-grey">
                    No players found
                  </q-item-section>
                </q-item>
              </template>
            </q-select>

            <q-btn
              color="primary"
              label="Add Player"
              @click="addPlayerToLeague"
              :disable="!selectedPlayer"
              :loading="loading"
              class="full-width"
            />
          </q-card-section>
        </q-card>
      </div>

      <!-- Quick Match Section -->
      <div class="col-12 col-md-6">
        <q-card>
          <q-card-section>
            <div class="text-h6 q-mb-md">
              <q-icon name="sports_score" class="q-mr-sm" />
              Quick Match
            </div>

            <q-select
              v-model="quickMatch.player1"
              :options="currentLeague?.players || []"
              option-label="name"
              option-value="id"
              label="Player 1"
              class="q-mb-md"
            />

            <q-select
              v-model="quickMatch.player2"
              :options="availableOpponents"
              option-label="name"
              option-value="id"
              label="Player 2"
              class="q-mb-md"
            />

            <q-btn
              color="primary"
              label="Create & Start Match"
              @click="createQuickMatch"
              :disable="
                !quickMatch.player1 ||
                !quickMatch.player2 ||
                quickMatch.player1.id === quickMatch.player2.id
              "
              :loading="loading"
              class="full-width q-mb-sm"
            />

            <q-btn
              color="secondary"
              label="Go to Matches"
              @click="goToMatches"
              outline
              class="full-width"
            />
          </q-card-section>
        </q-card>
      </div>

      <!-- Create Tournament Section -->
      <div class="col-12 col-md-6">
        <q-card>
          <q-card-section>
            <div class="text-h6 q-mb-md">
              <q-icon name="emoji_events" class="q-mr-sm" />
              Create Tournament
            </div>

            <q-input
              v-model="newTournament.name"
              label="Tournament Name"
              class="q-mb-md"
            />

            <q-select
              v-model="newTournament.format"
              :options="tournamentFormats"
              label="Tournament Format"
              class="q-mb-md"
            />

            <q-btn
              color="secondary"
              label="Create Tournament"
              @click="createTournament"
              :disable="!newTournament.name || !newTournament.format"
              :loading="loading"
              class="full-width"
            />
          </q-card-section>
        </q-card>
      </div>

      <!-- Current Players -->
      <div class="col-12 col-md-6">
        <q-card>
          <q-card-section>
            <div class="text-h6 q-mb-md">
              <q-icon name="group" class="q-mr-sm" />
              League Players
            </div>

            <q-list v-if="currentLeague?.players?.length" separator>
              <q-item v-for="player in currentLeague.players" :key="player.id">
                <q-item-section avatar>
                  <q-avatar color="primary" text-color="white">
                    {{ player.name.charAt(0).toUpperCase() }}
                  </q-avatar>
                </q-item-section>
                <q-item-section>
                  <q-item-label>{{ player.name }}</q-item-label>
                  <q-item-label caption>{{ player.email }}</q-item-label>
                </q-item-section>
              </q-item>
            </q-list>

            <div v-else class="text-center text-grey-5 q-pa-md">
              No players in this league yet
            </div>
          </q-card-section>
        </q-card>
      </div>

      <!-- League Tournaments -->
      <div class="col-12 col-md-6">
        <q-card>
          <q-card-section>
            <div class="text-h6 q-mb-md">
              <q-icon name="emoji_events" class="q-mr-sm" />
              League Tournaments
            </div>

            <q-list v-if="leagueTournaments?.length" separator>
              <q-item
                v-for="tournament in leagueTournaments"
                :key="tournament.id"
                clickable
                @click="goToTournament(tournament.id)"
              >
                <q-item-section>
                  <q-item-label>{{ tournament.name }}</q-item-label>
                  <q-item-label caption>
                    Status: {{ tournament.status }} | Format:
                    {{ tournament.format }}
                  </q-item-label>
                </q-item-section>
                <q-item-section side>
                  <q-chip
                    :color="tournament.status === 'active' ? 'green' : 'orange'"
                    text-color="white"
                    size="sm"
                  >
                    {{ tournament.status }}
                  </q-chip>
                </q-item-section>
              </q-item>
            </q-list>

            <div v-else class="text-center text-grey-5 q-pa-md">
              No tournaments in this league yet
            </div>
          </q-card-section>
        </q-card>
      </div>
      <q-btn @click="testPrint">Test</q-btn>
    </div>
  </q-page>
</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useQuasar } from "quasar";
import { useLeaguesStore } from "src/stores/leagues";
import { useTournamentsStore } from "src/stores/tournaments";
import { useMatchesStore } from "src/stores/matches";
import { playerApi } from "src/services/api";

const route = useRoute();
const router = useRouter();
const $q = useQuasar();
const leaguesStore = useLeaguesStore();
const tournamentsStore = useTournamentsStore();
const matchesStore = useMatchesStore();

const selectedPlayer = ref(null);
const allPlayers = ref([]);
const availablePlayers = ref([]);
const newTournament = ref({
  name: "",
  format: "",
});

const quickMatch = ref({
  player1: null,
  player2: null,
});

const tournamentFormats = [
  "Single Elimination",
  "Double Elimination",
  "Round Robin",
  "Swiss System",
];

const loading = computed(
  () => leaguesStore.loading || tournamentsStore.loading || matchesStore.loading
);
const currentLeague = computed(() => leaguesStore.currentLeague);
const leagueTournaments = computed(() => leaguesStore.leagueTournaments);

const availableOpponents = computed(() => {
  if (!quickMatch.value.player1 || !currentLeague.value?.players)
    return currentLeague.value?.players || [];
  return currentLeague.value.players.filter(
    (player) => player.id !== quickMatch.value.player1.id
  );
});

function testPrint() {
  console.log(currentLeague);
}

onMounted(async () => {
  const leagueId = route.params.id;
  await fetchLeagueData(leagueId);
  await fetchPlayers();
});

async function fetchLeagueData(leagueId) {
  try {
    await Promise.all([
      leaguesStore.fetchLeague(leagueId),
      leaguesStore.fetchLeagueTournaments(leagueId),
    ]);
  } catch (error) {
    $q.notify({
      type: "negative",
      message: "Failed to load league data",
    });
  }
}

async function fetchPlayers() {
  try {
    const response = await playerApi.getPlayers();
    console.log("asd players in league", response);
    allPlayers.value = Array.isArray(response.data.data)
      ? response.data.data
      : [];
    filterAvailablePlayers();
  } catch (error) {
    console.error("Error fetching players:", error);
  }
}

function filterAvailablePlayers() {
  const leaguePlayerIds = currentLeague.value?.players?.map((p) => p.id) || [];
  availablePlayers.value = allPlayers.value.filter(
    (player) => !leaguePlayerIds.includes(player.id)
  );
}

function filterPlayers(val, update) {
  update(() => {
    if (val === "") {
      filterAvailablePlayers();
    } else {
      const needle = val.toLowerCase();
      availablePlayers.value = allPlayers.value.filter((player) => {
        const leaguePlayerIds =
          currentLeague.value?.players?.map((p) => p.id) || [];
        return (
          !leaguePlayerIds.includes(player.id) &&
          player.name.toLowerCase().includes(needle)
        );
      });
    }
  });
}

async function addPlayerToLeague() {
  if (!selectedPlayer.value) return;

  try {
    await leaguesStore.addPlayerToLeague(
      route.params.id,
      selectedPlayer.value.id
    );

    $q.notify({
      type: "positive",
      message: `${selectedPlayer.value.name} added to league successfully`,
    });

    selectedPlayer.value = null;
    filterAvailablePlayers();
  } catch (error) {
    $q.notify({
      type: "negative",
      message: "Failed to add player to league",
    });
  }
}

async function createTournament() {
  if (!newTournament.value.name || !newTournament.value.format) return;

  try {
    const tournamentData = {
      name: newTournament.value.name,
      format: newTournament.value.format,
      league_id: parseInt(route.params.id),
    };

    await tournamentsStore.createTournament(tournamentData);

    $q.notify({
      type: "positive",
      message: "Tournament created successfully",
    });

    newTournament.value = { name: "", format: "" };
    await leaguesStore.fetchLeagueTournaments(route.params.id);
  } catch (error) {
    $q.notify({
      type: "negative",
      message: "Failed to create tournament",
    });
  }
}

async function startLeague() {
  try {
    await leaguesStore.startLeague(route.params.id);

    $q.notify({
      type: "positive",
      message: "League started successfully",
    });
  } catch (error) {
    $q.notify({
      type: "negative",
      message: "Failed to start league",
    });
  }
}

function goToTournament(tournamentId) {
  router.push(`/tournaments/${tournamentId}`);
}

async function createQuickMatch() {
  if (!quickMatch.value.player1 || !quickMatch.value.player2) return;

  try {
    const matchData = {
      player1_id: quickMatch.value.player1.id,
      player2_id: quickMatch.value.player2.id,
    };

    const createdMatch = await matchesStore.createMatch(matchData);

    // Start the match immediately
    await matchesStore.startMatch(createdMatch.id);

    $q.notify({
      type: "positive",
      message: "Match created and started successfully",
    });

    // Navigate to the match
    router.push(`/matches/${createdMatch.id}`);
  } catch (error) {
    $q.notify({
      type: "negative",
      message: "Failed to create match",
    });
  }
}

function goToMatches() {
  router.push("/matches");
}
</script>
