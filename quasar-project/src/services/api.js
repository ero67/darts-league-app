import axios from 'axios'

const API_BASE_URL = 'http://localhost:8080/api'

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
})

// Request interceptor for adding auth token if needed
api.interceptors.request.use(
  (config) => {
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// Response interceptor for handling errors
api.interceptors.response.use(
  (response) => {
    return response
  },
  (error) => {
    console.error('API Error:', error)
    return Promise.reject(error)
  }
)

// Player API
export const playerApi = {
  getPlayers: () => api.get('/players'),
  createPlayer: (player) => api.post('/players', player),
  getPlayer: (id) => api.get(`/players/${id}`),
  updatePlayer: (id, player) => api.put(`/players/${id}`, player),
  deletePlayer: (id) => api.delete(`/players/${id}`),
  searchPlayers: (query) => api.get(`/players/search?q=${query}`),
  getPlayerMatches: (id) => api.get(`/players/${id}/matches`),
}

// League API
export const leagueApi = {
  getLeagues: () => api.get('/leagues'),
  createLeague: (league) => api.post('/leagues', league),
  getLeague: (id) => api.get(`/leagues/${id}`),
  addPlayerToLeague: (id, playerId) => api.post(`/leagues/${id}/players`, { player_id: playerId }),
  getLeagueStandings: (id) => api.get(`/leagues/${id}/standings`),
  startLeague: (id) => api.post(`/leagues/${id}/start`),
  getLeagueTournaments: (id) => api.get(`/leagues/${id}/tournaments`),
}

// Tournament API
export const tournamentApi = {
  createTournament: (tournament) => api.post('/tournaments', tournament),
  getTournament: (id) => api.get(`/tournaments/${id}`),
  addPlayerToTournament: (id, playerId) => api.post(`/tournaments/${id}/players`, { player_id: playerId }),
  startTournament: (id) => api.post(`/tournaments/${id}/start`),
  getTournamentMatches: (id) => api.get(`/tournaments/${id}/matches`),
}

// Match API
export const matchApi = {
  createMatch: (match) => api.post('/matches', match),
  getMatch: (id) => api.get(`/matches/${id}`),
  startMatch: (id) => api.post(`/matches/${id}/start`),
  updateMatchScore: (id, score) => api.put(`/matches/${id}/score`, score),
  completeMatch: (id, winnerId) => api.post(`/matches/${id}/complete`, { winner_id: winnerId }),
}

// Health check - uses different base URL (no /api prefix)
const healthClient = axios.create({
  baseURL: 'http://localhost:8080',
  headers: {
    'Content-Type': 'application/json',
  },
})

export const healthApi = {
  check: () => healthClient.get('/health'),
}

export default api