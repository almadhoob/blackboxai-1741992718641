<template>
  <div class="min-h-screen bg-gray-50">
    <div class="max-w-7xl mx-auto py-12 px-4 sm:px-6 lg:px-8">
      <!-- Profile header -->
      <div class="bg-white shadow rounded-lg overflow-hidden">
        <div class="relative h-48 bg-primary-600">
          <!-- Cover image -->
          <img
            class="absolute inset-0 h-full w-full object-cover"
            src="https://images.unsplash.com/photo-1522071820081-009f0129c71c?ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80"
            alt="Cover"
          />
        </div>
        <div class="relative -mt-16 px-6 pb-6">
          <!-- Avatar -->
          <div class="flex items-end space-x-5">
            <div class="flex-shrink-0">
              <img
                class="h-32 w-32 rounded-full ring-4 ring-white bg-white"
                :src="profile.avatar"
                :alt="profile.name"
              />
            </div>
            <div class="flex-1 min-w-0 flex items-center justify-between space-x-6 pb-1">
              <div>
                <h1 class="text-2xl font-bold text-gray-900 truncate">{{ profile.name }}</h1>
                <p class="text-sm text-gray-500">{{ profile.title }}</p>
              </div>
              <div class="flex flex-col sm:flex-row sm:flex-wrap space-y-3 sm:space-y-0 sm:space-x-3">
                <button
                  type="button"
                  class="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
                >
                  <i class="fas fa-user-plus -ml-1 mr-2 h-5 w-5"></i>
                  Follow
                </button>
                <button
                  type="button"
                  class="inline-flex items-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
                >
                  <i class="fas fa-envelope -ml-1 mr-2 h-5 w-5"></i>
                  Message
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Profile content -->
      <div class="mt-6 grid grid-cols-1 gap-6 lg:grid-cols-3">
        <!-- Left sidebar -->
        <div class="lg:col-span-1">
          <div class="bg-white shadow rounded-lg">
            <div class="p-6">
              <h2 class="text-lg font-medium text-gray-900">About</h2>
              <p class="mt-4 text-sm text-gray-500">
                {{ profile.bio }}
              </p>
              <div class="mt-6">
                <div class="flex items-center space-x-2 text-sm text-gray-500">
                  <i class="fas fa-map-marker-alt"></i>
                  <span>{{ profile.location }}</span>
                </div>
                <div class="mt-3 flex items-center space-x-2 text-sm text-gray-500">
                  <i class="fas fa-link"></i>
                  <a :href="profile.website" class="text-primary-600 hover:text-primary-700">
                    {{ profile.website }}
                  </a>
                </div>
                <div class="mt-3 flex items-center space-x-2 text-sm text-gray-500">
                  <i class="fab fa-github"></i>
                  <a :href="profile.github" class="text-primary-600 hover:text-primary-700">
                    {{ profile.github }}
                  </a>
                </div>
              </div>
              <div class="mt-6">
                <h3 class="text-sm font-medium text-gray-900">Skills</h3>
                <div class="mt-2 flex flex-wrap gap-2">
                  <span
                    v-for="skill in profile.skills"
                    :key="skill"
                    class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-primary-100 text-primary-800"
                  >
                    {{ skill }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Main content -->
        <div class="lg:col-span-2">
          <!-- Activity feed -->
          <div class="bg-white shadow rounded-lg">
            <div class="p-6">
              <h2 class="text-lg font-medium text-gray-900">Activity</h2>
              <div class="flow-root mt-6">
                <ul role="list" class="-mb-8">
                  <li v-for="(activity, index) in profile.activities" :key="index">
                    <div class="relative pb-8">
                      <span
                        v-if="index !== profile.activities.length - 1"
                        class="absolute top-4 left-4 -ml-px h-full w-0.5 bg-gray-200"
                        aria-hidden="true"
                      ></span>
                      <div class="relative flex space-x-3">
                        <div>
                          <span
                            :class="[
                              activity.type === 'commit' ? 'bg-green-500' : 'bg-primary-500',
                              'h-8 w-8 rounded-full flex items-center justify-center ring-8 ring-white'
                            ]"
                          >
                            <i
                              :class="[
                                activity.type === 'commit' ? 'fas fa-code-branch' : 'fas fa-star',
                                'text-white'
                              ]"
                            ></i>
                          </span>
                        </div>
                        <div class="min-w-0 flex-1 flex justify-between space-x-4">
                          <div>
                            <p class="text-sm text-gray-500">
                              {{ activity.description }}
                              <a href="#" class="font-medium text-gray-900">
                                {{ activity.project }}
                              </a>
                            </p>
                          </div>
                          <div class="text-right text-sm whitespace-nowrap text-gray-500">
                            <time :datetime="activity.date">{{ activity.date }}</time>
                          </div>
                        </div>
                      </div>
                    </div>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// Sample profile data
const profile = {
  name: 'Sarah Johnson',
  title: 'Senior Software Engineer',
  avatar: 'https://via.placeholder.com/150',
  bio: 'Full-stack developer passionate about open source and distributed systems. Currently working on cloud-native applications and contributing to various open source projects.',
  location: 'San Francisco, CA',
  website: 'https://sarahjohnson.dev',
  github: 'https://github.com/sarahj',
  skills: [
    'Go',
    'Vue.js',
    'TypeScript',
    'PostgreSQL',
    'Docker',
    'Kubernetes',
    'GraphQL',
    'REST APIs'
  ],
  activities: [
    {
      type: 'commit',
      description: 'Pushed 3 commits to',
      project: 'kubernetes/kubernetes',
      date: '3 hours ago'
    },
    {
      type: 'star',
      description: 'Starred',
      project: 'vuejs/core',
      date: '1 day ago'
    },
    {
      type: 'commit',
      description: 'Created pull request in',
      project: 'golang/go',
      date: '2 days ago'
    },
    {
      type: 'star',
      description: 'Starred',
      project: 'tailwindlabs/tailwindcss',
      date: '4 days ago'
    }
  ]
}
</script>
