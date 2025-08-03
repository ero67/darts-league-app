import { defineStore } from "pinia";
import api from "../services/api";

export const useDashboardStore = defineStore("dashboard", {
  state: () => ({
    stats: {
      totalLeagues: 0,
      totalPlayers: 0,
      activeGames: 0,
      totalMatches: 0,
      recentMatches: [],
    },
    leagues: [],
    topPlayers: [],
    loading: false,
    error: null,
    lastUpdated: null,
  }),

  getters: {
    getActiveLeagues: (state) => {
      return state.leagues.filter((league) => league.status === "active");
    },

    getRecentActivity: (state) => {
      return state.stats.recentMatches.slice(0, 5);
    },
  },

  actions: {
    async fetchDashboardData() {
      this.loading = true;
      this.error = null;

      try {
        // Mock data for now - replace with actual API calls
        await new Promise((resolve) => setTimeout(resolve, 1000)); // Simulate API delay

        this.stats = {
          totalLeagues: 5,
          totalPlayers: 24,
          activeGames: 3,
          totalMatches: 142,
          recentMatches: [
            {
              id: 1,
              title: "John vs Mike",
              description: "League Championship Final",
              type: "match",
              createdAt: new Date().toISOString(),
            },
            {
              id: 2,
              title: "New League Created",
              description: "Summer Tournament 2025",
              type: "league",
              createdAt: new Date(Date.now() - 86400000).toISOString(),
            },
          ],
        };

        this.leagues = [
          {
            id: 1,
            name: "Championship League",
            status: "active",
            playerCount: 12,
            matchCount: 45,
          },
          {
            id: 2,
            name: "Beginners League",
            status: "active",
            playerCount: 8,
            matchCount: 23,
          },
        ];

        this.topPlayers = [
          {
            id: 1,
            name: "John Doe",
            rating: 1850,
            gamesPlayed: 45,
            winRate: 78,
          },
          {
            id: 2,
            name: "Mike Smith",
            rating: 1720,
            gamesPlayed: 38,
            winRate: 65,
          },
          {
            id: 3,
            name: "Sarah Johnson",
            rating: 1680,
            gamesPlayed: 42,
            winRate: 58,
          },
        ];

        this.lastUpdated = new Date().toISOString();
      } catch (error) {
        this.error = error.message || "Failed to fetch dashboard data";
        console.error("Dashboard fetch error:", error);
      } finally {
        this.loading = false;
      }
    },

    async refreshStats() {
      try {
        // Mock refresh - replace with actual API call
        await new Promise((resolve) => setTimeout(resolve, 500));
        this.stats.totalMatches += 1;
        this.lastUpdated = new Date().toISOString();
      } catch (error) {
        this.error = "Failed to refresh stats";
      }
    },
  },
});
