import { defineStore } from 'pinia'
import { tournamentApi } from 'src/services/api'

export const useTournamentsStore = defineStore('tournaments', {
  state: () => ({
    tournaments: [],
    currentTournament: null,
    tournamentMatches: [],
    loading: false,
    error: null,
  }),

  getters: {
    getTournamentById: (state) => (id) => {
      return state.tournaments.find(tournament => tournament.id === id)
    },
    activeTournaments: (state) => {
      return state.tournaments.filter(tournament => tournament.status === 'active')
    },
    setupTournaments: (state) => {
      return state.tournaments.filter(tournament => tournament.status === 'setup')
    },
  },

  actions: {
    async createTournament(tournamentData) {
      this.loading = true
      this.error = null
      try {
        const response = await tournamentApi.createTournament(tournamentData)
        this.tournaments.push(response.data)
        return response.data
      } catch (error) {
        this.error = error.message
        console.error('Error creating tournament:', error)
        throw error
      } finally {
        this.loading = false
      }
    },

    async fetchTournament(id) {
      this.loading = true
      this.error = null
      try {
        const response = await tournamentApi.getTournament(id)
        this.currentTournament = response.data
        return response.data
      } catch (error) {
        this.error = error.message
        console.error('Error fetching tournament:', error)
        throw error
      } finally {
        this.loading = false
      }
    },

    async addPlayerToTournament(tournamentId, playerId) {
      this.loading = true
      this.error = null
      try {
        await tournamentApi.addPlayerToTournament(tournamentId, playerId)
        // Refresh current tournament if it's the one we're adding to
        if (this.currentTournament && this.currentTournament.id === tournamentId) {
          await this.fetchTournament(tournamentId)
        }
      } catch (error) {
        this.error = error.message
        console.error('Error adding player to tournament:', error)
        throw error
      } finally {
        this.loading = false
      }
    },

    async startTournament(id) {
      this.loading = true
      this.error = null
      try {
        await tournamentApi.startTournament(id)
        // Update the tournament in our list
        const index = this.tournaments.findIndex(t => t.id === id)
        if (index !== -1) {
          this.tournaments[index].status = 'active'
        }
        // Update current tournament if it's the one we're starting
        if (this.currentTournament && this.currentTournament.id === id) {
          this.currentTournament.status = 'active'
        }
      } catch (error) {
        this.error = error.message
        console.error('Error starting tournament:', error)
        throw error
      } finally {
        this.loading = false
      }
    },

    async fetchTournamentMatches(id) {
      this.loading = true
      this.error = null
      try {
        const response = await tournamentApi.getTournamentMatches(id)
        this.tournamentMatches = response.data
        return response.data
      } catch (error) {
        this.error = error.message
        console.error('Error fetching tournament matches:', error)
        throw error
      } finally {
        this.loading = false
      }
    },

    clearError() {
      this.error = null
    },
  },
})