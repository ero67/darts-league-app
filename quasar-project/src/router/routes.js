const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/IndexPage.vue') },
      { path: 'players', component: () => import('pages/PlayersPage.vue') },
      { path: 'leagues', component: () => import('pages/LeaguesPage.vue') },
      { path: 'tournaments', component: () => import('pages/TournamentsPage.vue') },
      { path: 'matches', component: () => import('pages/MatchesPage.vue') },
      { path: 'standings', component: () => import('pages/StandingsPage.vue') },
      { path: 'leagues/:id', component: () => import('pages/LeagueDetailPage.vue') },
      { path: 'leagues/:id/standings', component: () => import('pages/LeagueStandingsPage.vue') },
      { path: 'tournaments/:id', component: () => import('pages/TournamentDetailPage.vue') },
      { path: 'matches/:id', component: () => import('pages/MatchDetailPage.vue') }
    ]
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue')
  }
]

export default routes
