<template>
  <q-page padding>
    <div class="row q-col-gutter-md">
      <div class="col-12">
        <q-card>
          <q-card-section>
            <div class="row items-center q-col-gutter-md">
              <div class="col-auto">
                <q-icon name="people" size="lg" class="text-primary" />
              </div>
              <div class="col">
                <div class="text-h4">Players</div>
                <div class="text-subtitle1">Manage your darts players</div>
              </div>
              <div class="col-auto">
                <q-btn
                  color="primary"
                  label="Add Player"
                  icon="person_add"
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
                <q-input
                  v-model="searchQuery"
                  outlined
                  label="Search players"
                  placeholder="Search by name or nickname"
                  clearable
                  @input="onSearchInput"
                >
                  <template v-slot:prepend>
                    <q-icon name="search" />
                  </template>
                </q-input>
              </div>
              <div class="col-md-3 col-12">
                <q-btn
                  outline
                  color="primary"
                  label="Refresh"
                  icon="refresh"
                  @click="refreshPlayers"
                  :loading="playersStore.loading"
                />
              </div>
            </div>

            <q-table
              :rows="filteredPlayers"
              :columns="columns"
              row-key="id"
              :loading="playersStore.loading"
              :pagination="{ rowsPerPage: 10 }"
              class="players-table"
            >
              <template v-slot:body-cell-avatar="props">
                <q-td :props="props">
                  <q-avatar size="40px" color="primary" text-color="white">
                    <img
                      v-if="props.row.avatar_url"
                      :src="props.row.avatar_url"
                    />
                    <span v-else>{{
                      props.row.name.charAt(0).toUpperCase()
                    }}</span>
                  </q-avatar>
                </q-td>
              </template>

              <template v-slot:body-cell-display_name="props">
                <q-td :props="props">
                  <div class="text-weight-medium">{{ props.row.name }}</div>
                  <div
                    v-if="props.row.nickname"
                    class="text-caption text-grey-6"
                  >
                    "{{ props.row.nickname }}"
                  </div>
                </q-td>
              </template>

              <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                  <q-btn
                    flat
                    round
                    dense
                    icon="edit"
                    color="primary"
                    @click="editPlayer(props.row)"
                  />
                  <q-btn
                    flat
                    round
                    dense
                    icon="delete"
                    color="negative"
                    @click="confirmDeletePlayer(props.row)"
                  />
                </q-td>
              </template>

              <template v-slot:no-data>
                <div class="full-width row flex-center text-accent q-gutter-sm">
                  <q-icon size="2em" name="sentiment_dissatisfied" />
                  <span>No players found</span>
                </div>
              </template>
            </q-table>
          </q-card-section>
        </q-card>
      </div>
    </div>

    <!-- Create/Edit Player Dialog -->
    <q-dialog v-model="showCreateDialog" persistent>
      <q-card style="min-width: 400px">
        <q-card-section>
          <div class="text-h6">
            {{ editingPlayer ? "Edit Player" : "Add New Player" }}
          </div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <q-form @submit="savePlayer" class="q-gutter-md">
            <q-input
              v-model="playerForm.name"
              label="Full Name *"
              outlined
              :rules="[(val) => !!val || 'Name is required']"
              ref="nameInput"
            />

            <q-input
              v-model="playerForm.nickname"
              label="Nickname"
              outlined
              hint="Optional display name"
            />

            <q-input
              v-model="playerForm.email"
              label="Email"
              type="email"
              outlined
              hint="Optional email address"
            />

            <q-input
              v-model="playerForm.avatar_url"
              label="Avatar URL"
              outlined
              hint="Optional image URL"
            />

            <div class="row q-col-gutter-sm">
              <div class="col">
                <q-btn label="Cancel" color="grey" flat @click="closeDialog" />
              </div>
              <div class="col">
                <q-btn
                  label="Save"
                  color="primary"
                  type="submit"
                  :loading="playersStore.loading"
                />
              </div>
            </div>
          </q-form>
        </q-card-section>
      </q-card>
    </q-dialog>

    <!-- Delete Confirmation Dialog -->
    <q-dialog v-model="showDeleteDialog" persistent>
      <q-card>
        <q-card-section class="row items-center">
          <q-avatar icon="warning" color="negative" text-color="white" />
          <span class="q-ml-sm"
            >Are you sure you want to delete this player?</span
          >
        </q-card-section>

        <q-card-actions align="right">
          <q-btn
            flat
            label="Cancel"
            color="primary"
            @click="showDeleteDialog = false"
          />
          <q-btn
            flat
            label="Delete"
            color="negative"
            @click="deletePlayer"
            :loading="playersStore.loading"
          />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </q-page>
</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import { useQuasar } from "quasar";
import { usePlayersStore } from "src/stores/players";

const $q = useQuasar();
const playersStore = usePlayersStore();

const showCreateDialog = ref(false);
const showDeleteDialog = ref(false);
const editingPlayer = ref(null);
const playerToDelete = ref(null);
const searchQuery = ref("");

const playerForm = ref({
  name: "",
  nickname: "",
  email: "",
  avatar_url: "",
});

const columns = [
  {
    name: "avatar",
    label: "",
    field: "avatar_url",
    align: "center",
    style: "width: 60px",
  },
  {
    name: "display_name",
    label: "Name",
    field: "name",
    align: "left",
    sortable: true,
  },
  {
    name: "email",
    label: "Email",
    field: "email",
    align: "left",
    sortable: true,
    format: (val) => val || "N/A",
  },
  {
    name: "created_at",
    label: "Created",
    field: "created_at",
    align: "left",
    sortable: true,
    format: (val) => new Date(val).toLocaleDateString(),
  },
  {
    name: "actions",
    label: "Actions",
    field: "actions",
    align: "center",
    style: "width: 120px",
  },
];

const filteredPlayers = computed(() => {
  // Ensure players is always an array
  const players = playersStore.players;
  // return players;
  if (!searchQuery.value) {
    return players;
  }

  const query = searchQuery.value.toLowerCase();
  return players.filter(
    (player) =>
      player.name.toLowerCase().includes(query) ||
      (player.nickname && player.nickname.toLowerCase().includes(query)) ||
      (player.email && player.email.toLowerCase().includes(query))
  );
});

function onSearchInput() {
  // Debounce search if needed
}

function refreshPlayers() {
  playersStore.fetchPlayers();
}

function editPlayer(player) {
  editingPlayer.value = player;
  playerForm.value = {
    name: player.name,
    nickname: player.nickname || "",
    email: player.email || "",
    avatar_url: player.avatar_url || "",
  };
  showCreateDialog.value = true;
}

function confirmDeletePlayer(player) {
  playerToDelete.value = player;
  showDeleteDialog.value = true;
}

async function deletePlayer() {
  try {
    await playersStore.deletePlayer(playerToDelete.value.id);
    showDeleteDialog.value = false;
    playerToDelete.value = null;
    $q.notify({
      type: "positive",
      message: "Player deleted successfully",
    });
  } catch (error) {
    $q.notify({
      type: "negative",
      message: "Failed to delete player",
    });
  }
}

async function savePlayer() {
  try {
    const playerData = {
      name: playerForm.value.name,
      nickname: playerForm.value.nickname || null,
      email: playerForm.value.email || null,
      avatar_url: playerForm.value.avatar_url || null,
    };

    if (editingPlayer.value) {
      await playersStore.updatePlayer(editingPlayer.value.id, playerData);
      $q.notify({
        type: "positive",
        message: "Player updated successfully",
      });
    } else {
      await playersStore.createPlayer(playerData);
      $q.notify({
        type: "positive",
        message: "Player created successfully",
      });
    }

    closeDialog();
  } catch (error) {
    $q.notify({
      type: "negative",
      message: "Failed to save player",
    });
  }
}

function closeDialog() {
  showCreateDialog.value = false;
  editingPlayer.value = null;
  playerForm.value = {
    name: "",
    nickname: "",
    email: "",
    avatar_url: "",
  };
}

onMounted(() => {
  playersStore.fetchPlayers();
});
</script>

<style scoped>
.players-table {
  /* Custom table styling if needed */
}
</style>
