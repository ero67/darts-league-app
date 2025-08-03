import { defineStore } from "pinia";
import { playerApi } from "src/services/api";

export const usePlayersStore = defineStore("players", {
  state: () => ({
    players: [],
    currentPlayer: null,
    loading: false,
    error: null,
  }),

  getters: {
    getPlayerById: (state) => (id) => {
      return state.players.find((player) => player.id === id);
    },
  },

  actions: {
    async fetchPlayers() {
      this.loading = true;
      this.error = null;
      try {
        const response = await playerApi.getPlayers();
        this.players = response.data.data;
      } catch (error) {
        this.error = error.message;
        this.players = []; // Ensure players remains an array on error
        console.error("Error fetching players:", error);
      } finally {
        this.loading = false;
      }
    },

    async createPlayer(playerData) {
      this.loading = true;
      this.error = null;
      try {
        const response = await playerApi.createPlayer(playerData);
        this.players.push(response.data.data);
        return response.data.data;
      } catch (error) {
        this.error = error.message;
        console.error("Error creating player:", error);
        throw error;
      } finally {
        this.loading = false;
      }
    },

    async fetchPlayer(id) {
      this.loading = true;
      this.error = null;
      try {
        const response = await playerApi.getPlayer(id);
        this.currentPlayer = response.data;
        return response.data;
      } catch (error) {
        this.error = error.message;
        console.error("Error fetching player:", error);
        throw error;
      } finally {
        this.loading = false;
      }
    },

    async updatePlayer(id, playerData) {
      this.loading = true;
      this.error = null;
      try {
        const response = await playerApi.updatePlayer(id, playerData);
        const index = this.players.findIndex((p) => p.id === id);
        if (index !== -1) {
          this.players[index] = response.data;
        }
        return response.data;
      } catch (error) {
        this.error = error.message;
        console.error("Error updating player:", error);
        throw error;
      } finally {
        this.loading = false;
      }
    },

    async deletePlayer(id) {
      this.loading = true;
      this.error = null;
      try {
        await playerApi.deletePlayer(id);
        this.players = this.players.filter((p) => p.id !== id);
      } catch (error) {
        this.error = error.message;
        console.error("Error deleting player:", error);
        throw error;
      } finally {
        this.loading = false;
      }
    },

    async searchPlayers(query) {
      this.loading = true;
      this.error = null;
      try {
        const response = await playerApi.searchPlayers(query);
        return response.data;
      } catch (error) {
        this.error = error.message;
        console.error("Error searching players:", error);
        throw error;
      } finally {
        this.loading = false;
      }
    },

    clearError() {
      this.error = null;
    },
  },
});
