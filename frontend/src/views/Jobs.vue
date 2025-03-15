<template>
  <div class="min-h-screen bg-gray-50">
    <div class="max-w-7xl mx-auto py-12 px-4 sm:px-6 lg:px-8">
      <!-- Header -->
      <div class="lg:flex lg:items-center lg:justify-between">
        <div class="flex-1 min-w-0">
          <h2 class="text-3xl font-bold leading-7 text-gray-900 sm:text-4xl sm:truncate">
            Job Opportunities
          </h2>
          <p class="mt-1 text-lg text-gray-500">
            Find remote-friendly jobs at companies that value open source contributions
          </p>
        </div>
        <div class="mt-5 flex lg:mt-0 lg:ml-4">
          <span class="sm:ml-3">
            <button
              type="button"
              class="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
            >
              <i class="fas fa-filter -ml-1 mr-2 h-5 w-5"></i>
              Filter
            </button>
          </span>
        </div>
      </div>

      <!-- Search and filters -->
      <div class="mt-8">
        <div class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
          <div class="col-span-1 sm:col-span-2 lg:col-span-3">
            <label for="search" class="sr-only">Search jobs</label>
            <div class="relative">
              <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <i class="fas fa-search text-gray-400"></i>
              </div>
              <input
                type="text"
                name="search"
                id="search"
                class="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-500 focus:outline-none focus:placeholder-gray-400 focus:ring-1 focus:ring-primary-500 focus:border-primary-500 sm:text-sm"
                placeholder="Search for jobs..."
              />
            </div>
          </div>
        </div>
      </div>

      <!-- Job listings -->
      <div class="mt-8 grid gap-6 lg:grid-cols-2">
        <div v-for="job in jobs" :key="job.id" class="bg-white shadow rounded-lg p-6">
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="flex-shrink-0 h-12 w-12">
                <img :src="job.companyLogo" :alt="job.companyName" class="h-12 w-12 rounded-full" />
              </div>
              <div class="ml-4">
                <h3 class="text-lg font-medium text-gray-900">
                  {{ job.title }}
                </h3>
                <div class="text-sm text-gray-500">
                  {{ job.companyName }} • {{ job.location }}
                </div>
              </div>
            </div>
            <div class="flex items-center">
              <span
                :class="[
                  job.type === 'Full-time' ? 'bg-green-100 text-green-800' : 'bg-blue-100 text-blue-800',
                  'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium'
                ]"
              >
                {{ job.type }}
              </span>
            </div>
          </div>
          <div class="mt-4">
            <p class="text-sm text-gray-500">
              {{ job.description }}
            </p>
          </div>
          <div class="mt-4">
            <div class="flex space-x-2">
              <span
                v-for="tag in job.tags"
                :key="tag"
                class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-gray-100 text-gray-800"
              >
                {{ tag }}
              </span>
            </div>
          </div>
          <div class="mt-6 flex items-center justify-between">
            <div class="flex items-center">
              <i class="fas fa-clock text-gray-400"></i>
              <span class="ml-2 text-sm text-gray-500">Posted {{ job.postedAt }}</span>
            </div>
            <router-link
              :to="{ name: 'JobDetails', params: { id: job.id }}"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
            >
              View Details
            </router-link>
          </div>
        </div>
      </div>

      <!-- Pagination -->
      <div class="mt-8 flex justify-center">
        <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
          <a
            href="#"
            class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
          >
            <span class="sr-only">Previous</span>
            <i class="fas fa-chevron-left h-5 w-5"></i>
          </a>
          <a
            href="#"
            class="relative inline-flex items-center px-4 py-2 border border-gray-300 bg-white text-sm font-medium text-gray-700 hover:bg-gray-50"
          >
            1
          </a>
          <a
            href="#"
            class="relative inline-flex items-center px-4 py-2 border border-gray-300 bg-white text-sm font-medium text-gray-700 hover:bg-gray-50"
          >
            2
          </a>
          <a
            href="#"
            class="relative inline-flex items-center px-4 py-2 border border-gray-300 bg-white text-sm font-medium text-gray-700 hover:bg-gray-50"
          >
            3
          </a>
          <a
            href="#"
            class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
          >
            <span class="sr-only">Next</span>
            <i class="fas fa-chevron-right h-5 w-5"></i>
          </a>
        </nav>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// Sample job data
const jobs = [
  {
    id: 1,
    title: 'Senior Frontend Developer',
    companyName: 'TechCorp',
    companyLogo: 'https://via.placeholder.com/150',
    location: 'Remote',
    type: 'Full-time',
    description: 'We are looking for a senior frontend developer with experience in Vue.js and TypeScript...',
    tags: ['Vue.js', 'TypeScript', 'Tailwind CSS'],
    postedAt: '2 days ago'
  },
  {
    id: 2,
    title: 'Backend Engineer',
    companyName: 'OpenSource Inc',
    companyLogo: 'https://via.placeholder.com/150',
    location: 'San Francisco, CA',
    type: 'Contract',
    description: 'Join our team as a backend engineer working on distributed systems...',
    tags: ['Go', 'PostgreSQL', 'Docker'],
    postedAt: '1 week ago'
  }
]
</script>
