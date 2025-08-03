import { defineStore } from 'pinia'
import { matchApi } from 'src/services/api'

export const useMatchesStore = defineStore('matches', {
  state: () => ({
    matches: [],
    currentMatch: null,
    loading: false,
    error: null,
  }),

  getters: {
    getMatchById: (state) => (id) => {
      return state.matches.find(match => match.id === id)
    },
    activeMatches: (state) => {
      return state.matches.filter(match => match.status === 'in_progress')
    },
    pendingMatches: (state) => {
      return state.matches.filter(match => match.status === 'pending')
    },
    completedMatches: (state) => {
      return state.matches.filter(match => match.status === 'completed')
    },
  },

  actions: {
    async createMatch(matchData) {
      this.loading = true
      this.error = null
      try {
        const response = await matchApi.createMatch(matchData)
        this.matches.push(response.data)
        return response.data
      } catch (error) {
        this.error = error.message
        console.error('Error creating match:', error)
        throw error
      } finally {
        this.loading = false
      }
    },

    async fetchMatch(id) {
      this.loading = true
      this.error = null
      try {
        const response = await matchApi.getMatch(id)
        this.currentMatch = response.data
        return response.data
      } catch (error) {
        this.error = error.message
        console.error('Error fetching match:', error)
        throw error
      } finally {
        this.loading = false
      }
    },

    async startMatch(id) {
      this.loading = true
      this.error = null
      try {
        await matchApi.startMatch(id)
        // Update the match in our list
        const index = this.matches.findIndex(m => m.id === id)
        if (index !== -1) {
          this.matches[index].status = 'in_progress'
        }
        // Update current match if it's the one we're starting
        if (this.currentMatch && this.currentMatch.id === id) {
          this.currentMatch.status = 'in_progress'
        }
      } catch (error) {
        this.error = error.message
        console.error('Error starting match:', error)
        throw error
      } finally {
        this.loading = false
      }
    },

    async updateMatchScore(id, scoreData) {
      this.loading = true
      this.error = null
      try {
        await matchApi.updateMatchScore(id, scoreData)
        // Update the match in our list
        const index = this.matches.findIndex(m => m.id === id)
        if (index !== -1) {
          this.matches[index] = { ...this.matches[index], ...scoreData }
        }
        // Update current match if it's the one we're updating
        if (this.currentMatch && this.currentMatch.id === id) {
          this.currentMatch = { ...this.currentMatch, ...scoreData }
        }
      } catch (error) {
        this.error = error.message
        console.error('Error updating match score:', error)
        throw error
      } finally {
        this.loading = false
      }
    },

    async completeMatch(id, winnerId) {
      this.loading = true
      this.error = null
      try {
        await matchApi.completeMatch(id, winnerId)
        // Update the match in our list
        const index = this.matches.findIndex(m => m.id === id)
        if (index !== -1) {
          this.matches[index].status = 'completed'
          this.matches[index].winner_id = winnerId
        }
        // Update current match if it's the one we're completing
        if (this.currentMatch && this.currentMatch.id === id) {
          this.currentMatch.status = 'completed'
          this.currentMatch.winner_id = winnerId
        }
      } catch (error) {
        this.error = error.message
        console.error('Error completing match:', error)
        throw error
      } finally {
        this.loading = false
      }
    },

    setMatches(matches) {
      this.matches = matches
    },

    clearError() {
      this.error = null
    },
  },
})