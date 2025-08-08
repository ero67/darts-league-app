import { defineStore } from "pinia";
import { leagueApi } from "src/services/api";

export const useLeaguesStore = defineStore("leagues", {
  state: () => ({
    leagues: [],
    currentLeague: null,
    leagueStandings: [],
    leagueTournaments: [],
    loading: false,
    error: null,
  }),

  getters: {
    getLeagueById: (state) => (id) => {
      return state.leagues.find((league) => league.id === id);
    },
    activeLeagues: (state) => {
      return state.leagues.filter((league) => league.status === "active");
    },
    setupLeagues: (state) => {
      return state.leagues.filter((league) => league.status === "setup");
    },
  },

  actions: {
    async fetchLeagues() {
      this.loading = true;
      this.error = null;
      try {
        const response = await leagueApi.getLeagues();
        this.leagues = Array.isArray(response.data.data)
          ? response.data.data
          : [];
      } catch (error) {
        this.error = error.message;
        this.leagues = []; // Ensure leagues remains an array on error
        console.error("Error fetching leagues:", error);
      } finally {
        this.loading = false;
      }
    },

    async createLeague(leagueData) {
      this.loading = true;
      this.error = null;
      try {
        const response = await leagueApi.createLeague(leagueData);
        this.leagues.push(response.data);
        return response.data;
      } catch (error) {
        this.error = error.message;
        console.error("Error creating league:", error);
        throw error;
      } finally {
        this.loading = false;
      }
    },
    async fetchLeague(id) {
      this.loading = true;
      this.error = null;
      try {
        const response = await leagueApi.getLeague(id);
        // Fix: Use response.data.data if API wraps the league in a 'data' property
        this.currentLeague = response.data.data || response.data;
        return this.currentLeague;
      } catch (error) {
        this.error = error.message;
        console.error("Error fetching league:", error);
        throw error;
      } finally {
        this.loading = false;
      }
    },

    async addPlayerToLeague(leagueId, playerId) {
      this.loading = true;
      this.error = null;
      try {
        await leagueApi.addPlayerToLeague(leagueId, playerId);
        // Refresh current league if it's the one we're adding to
        if (this.currentLeague && this.currentLeague.id === leagueId) {
          await this.fetchLeague(leagueId);
        }
      } catch (error) {
        this.error = error.message;
        console.error("Error adding player to league:", error);
        throw error;
      } finally {
        this.loading = false;
      }
    },

    async fetchLeagueStandings(id) {
      this.loading = true;
      this.error = null;
      try {
        const response = await leagueApi.getLeagueStandings(id);
        this.leagueStandings = response.data;
        return response.data;
      } catch (error) {
        this.error = error.message;
        console.error("Error fetching league standings:", error);
        throw error;
      } finally {
        this.loading = false;
      }
    },

    async startLeague(id) {
      this.loading = true;
      this.error = null;
      try {
        await leagueApi.startLeague(id);
        // Update the league in our list
        const index = this.leagues.findIndex((l) => l.id === id);
        if (index !== -1) {
          this.leagues[index].status = "active";
        }
        // Update current league if it's the one we're starting
        if (this.currentLeague && this.currentLeague.id === id) {
          this.currentLeague.status = "active";
        }
      } catch (error) {
        this.error = error.message;
        console.error("Error starting league:", error);
        throw error;
      } finally {
        this.loading = false;
      }
    },

    async fetchLeagueTournaments(id) {
      this.loading = true;
      this.error = null;
      try {
        const response = await leagueApi.getLeagueTournaments(id);
        this.leagueTournaments = response.data;
        return response.data;
      } catch (error) {
        this.error = error.message;
        console.error("Error fetching league tournaments:", error);
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
