<template>
  <q-page padding>
    <div class="row q-col-gutter-md">
      <div class="col-12">
        <q-card>
          <q-card-section>
            <div class="row items-center q-col-gutter-md">
              <div class="col-auto">
                <q-icon name="emoji_events" size="lg" class="text-primary" />
              </div>
              <div class="col">
                <div class="text-h4">Leagues</div>
                <div class="text-subtitle1">Manage your darts leagues</div>
              </div>
              <div class="col-auto">
                <q-btn
                  color="primary"
                  label="Create League"
                  icon="add_circle"
                  @click="showCreateDialog = true"
                />
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>

      <div class="col-12">
        <q-card>
          <q-card-section>
            <div class="row q-col-gutter-md q-mb-md">
              <div class="col-md-6 col-12">
                <q-select
                  v-model="statusFilter"
                  :options="statusOptions"
                  label="Filter by Status"
                  outlined
                  clearable
                  emit-value
                  map-options
                />
              </div>
              <div class="col-md-3 col-12">
                <q-btn
                  outline
                  color="primary"
                  label="Refresh"
                  icon="refresh"
                  @click="refreshLeagues"
                  :loading="leaguesStore.loading"
                />
              </div>
            </div>

            <div class="row q-col-gutter-md">
              <div
                v-for="league in filteredLeagues"
                :key="league.id"
                class="col-12 col-md-6 col-lg-4"
              >
                <q-card class="league-card">
                  <q-card-section>
                    <div class="row items-center q-col-gutter-sm">
                      <div class="col">
                        <div class="text-h6">{{ league.name }}</div>
                        <div
                          v-if="league.description"
                          class="text-caption text-grey-6"
                        >
                          {{ league.description }}
                        </div>
                      </div>
                      <div class="col-auto">
                        <q-chip
                          :color="getStatusColor(league.status)"
                          text-color="white"
                          :label="league.status"
                        />
                      </div>
                    </div>
                  </q-card-section>

                  <q-card-section class="q-pt-none">
                    <div class="row q-col-gutter-sm text-caption">
                      <div class="col-6">
                        <q-icon name="calendar_today" size="xs" />
                        Season: {{ league.season || "N/A" }}
                      </div>
                      <div class="col-6">
                        <q-icon name="people" size="xs" />
                        Max Players: {{ league.max_players || "Unlimited" }}
                      </div>
                    </div>

                    <q-separator class="q-my-md" />

                    <div class="row q-col-gutter-xs text-caption">
                      <div class="col-4">
                        <div class="text-weight-medium">
                          {{ league.points_for_win }}
                        </div>
                        <div>Win Points</div>
                      </div>
                      <div class="col-4">
                        <div class="text-weight-medium">
                          {{ league.points_for_runner_up }}
                        </div>
                        <div>Runner-up</div>
                      </div>
                      <div class="col-4">
                        <div class="text-weight-medium">
                          {{ league.points_for_semi_final }}
                        </div>
                        <div>Semi-final</div>
                      </div>
                    </div>
                  </q-card-section>

                  <q-card-actions align="right">
                    <q-btn
                      flat
                      color="primary"
                      label="View"
                      @click="viewLeague(league)"
                    />
                    <q-btn
                      v-if="league.status === 'setup'"
                      flat
                      color="positive"
                      label="Start"
                      @click="startLeague(league)"
                    />
                    <q-btn-dropdown
                      flat
                      color="grey"
                      icon="more_vert"
                      dropdown-icon="none"
                    >
                      <q-list>
                        <q-item
                          clickable
                          v-close-popup
                          @click="editLeague(league)"
                        >
                          <q-item-section avatar>
                            <q-icon name="edit" />
                          </q-item-section>
                          <q-item-section>Edit</q-item-section>
                        </q-item>
                        <q-item
                          clickable
                          v-close-popup
                          @click="viewStandings(league)"
                        >
                          <q-item-section avatar>
                            <q-icon name="leaderboard" />
                          </q-item-section>
                          <q-item-section>Standings</q-item-section>
                        </q-item>
                        <q-item
                          clickable
                          v-close-popup
                          @click="addPlayersToLeague(league)"
                        >
                          <q-item-section avatar>
                            <q-icon name="person_add" />
                          </q-item-section>
                          <q-item-section>Add Players</q-item-section>
                        </q-item>
                      </q-list>
                    </q-btn-dropdown>
                  </q-card-actions>
                </q-card>
              </div>
            </div>

            <div
              v-if="filteredLeagues.length === 0"
              class="text-center q-pa-lg"
            >
              <q-icon name="emoji_events" size="3rem" class="text-grey-5" />
              <div class="text-h6 text-grey-5 q-mt-md">No leagues found</div>
              <div class="text-body2 text-grey-5">
                Create your first league to get started
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>
    </div>

    <!-- Create/Edit League Dialog -->
    <q-dialog v-model="showCreateDialog" persistent>
      <q-card style="min-width: 500px">
        <q-card-section>
          <div class="text-h6">
            {{ editingLeague ? "Edit League" : "Create New League" }}
          </div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <q-form @submit="saveLeague" class="q-gutter-md">
            <q-input
              v-model="leagueForm.name"
              label="League Name *"
              outlined
              :rules="[(val) => !!val || 'League name is required']"
            />

            <q-input
              v-model="leagueForm.description"
              label="Description"
              outlined
              type="textarea"
              rows="2"
              hint="Optional description"
            />

            <q-input
              v-model="leagueForm.season"
              label="Season"
              outlined
              hint="e.g., 2024 Spring, Summer 2024"
            />

            <q-input
              v-model.number="leagueForm.max_players"
              label="Maximum Players"
              outlined
              type="number"
              min="2"
              hint="Leave empty for unlimited"
            />

            <q-separator />

            <div class="text-subtitle2">Points System</div>

            <div class="row q-col-gutter-md">
              <div class="col-4">
                <q-input
                  v-model.number="leagueForm.points_for_win"
                  label="Win Points"
                  outlined
                  type="number"
                  min="0"
                />
              </div>
              <div class="col-4">
                <q-input
                  v-model.number="leagueForm.points_for_runner_up"
                  label="Runner-up Points"
                  outlined
                  type="number"
                  min="0"
                />
              </div>
              <div class="col-4">
                <q-input
                  v-model.number="leagueForm.points_for_semi_final"
                  label="Semi-final Points"
                  outlined
                  type="number"
                  min="0"
                />
              </div>
            </div>

            <div class="row q-col-gutter-sm">
              <div class="col">
                <q-btn label="Cancel" color="grey" flat @click="closeDialog" />
              </div>
              <div class="col">
                <q-btn
                  label="Save"
                  color="primary"
                  type="submit"
                  :loading="leaguesStore.loading"
                />
              </div>
            </div>
          </q-form>
        </q-card-section>
      </q-card>
    </q-dialog>

    <!-- Add Players Dialog -->
    <q-dialog v-model="showAddPlayersDialog" persistent>
      <q-card style="min-width: 400px">
        <q-card-section>
          <div class="text-h6">Add Players to {{ selectedLeague?.name }}</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <q-select
            v-model="selectedPlayers"
            :options="availablePlayers"
            label="Select Players"
            outlined
            multiple
            use-chips
            option-value="id"
            option-label="name"
            emit-value
            map-options
            :loading="playersStore.loading"
          />

          <div class="row q-col-gutter-sm q-mt-md">
            <div class="col">
              <q-btn
                label="Cancel"
                color="grey"
                flat
                @click="showAddPlayersDialog = false"
              />
            </div>
            <div class="col">
              <q-btn
                label="Add Players"
                color="primary"
                @click="addPlayersToLeagueTemp"
                :loading="leaguesStore.loading"
                :disable="selectedPlayers.length === 0"
              />
            </div>
          </div>
        </q-card-section>
      </q-card>
    </q-dialog>
  </q-page>
</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import { useQuasar } from "quasar";
import { useRouter } from "vue-router";
import { useLeaguesStore } from "src/stores/leagues";
import { usePlayersStore } from "src/stores/players";

const $q = useQuasar();
const router = useRouter();
const leaguesStore = useLeaguesStore();
const playersStore = usePlayersStore();

const showCreateDialog = ref(false);
const showAddPlayersDialog = ref(false);
const editingLeague = ref(null);
const selectedLeague = ref(null);
const selectedPlayers = ref([]);
const statusFilter = ref("");

const statusOptions = [
  { label: "All", value: "" },
  { label: "Setup", value: "setup" },
  { label: "Active", value: "active" },
  { label: "Completed", value: "completed" },
];

const leagueForm = ref({
  name: "",
  description: "",
  season: "",
  max_players: null,
  points_for_win: 3,
  points_for_runner_up: 2,
  points_for_semi_final: 1,
});

const filteredLeagues = computed(() => {
  if (!statusFilter.value) {
    return leaguesStore.leagues;
  }
  return leaguesStore.leagues.filter(
    (league) => league.status === statusFilter.value
  );
});

const availablePlayers = computed(() => {
  return playersStore.players.map((player) => ({
    id: player.id,
    name: player.nickname || player.name,
  }));
});

function getStatusColor(status) {
  switch (status) {
    case "setup":
      return "orange";
    case "active":
      return "positive";
    case "completed":
      return "grey";
    default:
      return "grey";
  }
}

function refreshLeagues() {
  leaguesStore.fetchLeagues();
}

function viewLeague(league) {
  router.push(`/leagues/${league.id}`);
}

function editLeague(league) {
  editingLeague.value = league;
  leagueForm.value = {
    name: league.name,
    description: league.description || "",
    season: league.season || "",
    max_players: league.max_players,
    points_for_win: league.points_for_win,
    points_for_runner_up: league.points_for_runner_up,
    points_for_semi_final: league.points_for_semi_final,
  };
  showCreateDialog.value = true;
}

function viewStandings(league) {
  router.push(`/leagues/${league.id}/standings`);
}

function addPlayersToLeagueTemp(league) {
  selectedLeague.value = league;
  selectedPlayers.value = [];
  showAddPlayersDialog.value = true;
}

async function startLeague(league) {
  try {
    await leaguesStore.startLeague(league.id);
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

async function saveLeague() {
  try {
    const leagueData = {
      name: leagueForm.value.name,
      description: leagueForm.value.description || null,
      season: leagueForm.value.season || null,
      max_players: leagueForm.value.max_players || null,
      points_for_win: leagueForm.value.points_for_win,
      points_for_runner_up: leagueForm.value.points_for_runner_up,
      points_for_semi_final: leagueForm.value.points_for_semi_final,
    };

    if (editingLeague.value) {
      // Update league functionality would need to be implemented in the backend
      $q.notify({
        type: "info",
        message: "League update functionality not yet implemented",
      });
    } else {
      await leaguesStore.createLeague(leagueData);
      $q.notify({
        type: "positive",
        message: "League created successfully",
      });
    }

    closeDialog();
  } catch (error) {
    $q.notify({
      type: "negative",
      message: "Failed to save league",
    });
  }
}

async function addPlayersToLeague() {
  try {
    for (const playerId of selectedPlayers.value) {
      await leaguesStore.addPlayerToLeague(selectedLeague.value.id, playerId);
    }

    showAddPlayersDialog.value = false;
    selectedPlayers.value = [];

    $q.notify({
      type: "positive",
      message: "Players added to league successfully",
    });
  } catch (error) {
    $q.notify({
      type: "negative",
      message: "Failed to add players to league",
    });
  }
}

function closeDialog() {
  showCreateDialog.value = false;
  editingLeague.value = null;
  leagueForm.value = {
    name: "",
    description: "",
    season: "",
    max_players: null,
    points_for_win: 3,
    points_for_runner_up: 2,
    points_for_semi_final: 1,
  };
}

onMounted(() => {
  leaguesStore.fetchLeagues();
  playersStore.fetchPlayers();
});
</script>

<style scoped>
.league-card {
  transition: transform 0.2s;
}

.league-card:hover {
  transform: translateY(-2px);
}
</style>
