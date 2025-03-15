<template>
  <div class="min-h-screen bg-gray-50">
    <div class="max-w-7xl mx-auto py-12 px-4 sm:px-6 lg:px-8">
      <!-- Back button -->
      <div class="mb-8">
        <router-link
          to="/jobs"
          class="inline-flex items-center text-sm font-medium text-primary-600 hover:text-primary-700"
        >
          <i class="fas fa-arrow-left mr-2"></i>
          Back to Jobs
        </router-link>
      </div>

      <!-- Job details -->
      <div class="bg-white shadow rounded-lg overflow-hidden">
        <div class="px-6 py-8">
          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <div class="flex-shrink-0 h-16 w-16">
                <img :src="job.companyLogo" :alt="job.companyName" class="h-16 w-16 rounded-lg" />
              </div>
              <div class="ml-6">
                <h1 class="text-3xl font-bold text-gray-900">{{ job.title }}</h1>
                <div class="mt-2 flex items-center text-gray-500">
                  <span>{{ job.companyName }}</span>
                  <span class="mx-2">&bull;</span>
                  <span>{{ job.location }}</span>
                </div>
              </div>
            </div>
            <div>
              <span
                :class="[
                  job.type === 'Full-time' ? 'bg-green-100 text-green-800' : 'bg-blue-100 text-blue-800',
                  'inline-flex items-center px-3 py-1 rounded-full text-sm font-medium'
                ]"
              >
                {{ job.type }}
              </span>
            </div>
          </div>

          <!-- Quick info -->
          <div class="mt-8 grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-4">
            <div class="border rounded-lg p-4">
              <div class="text-sm font-medium text-gray-500">Salary Range</div>
              <div class="mt-1 text-lg font-semibold text-gray-900">{{ job.salary }}</div>
            </div>
            <div class="border rounded-lg p-4">
              <div class="text-sm font-medium text-gray-500">Experience Level</div>
              <div class="mt-1 text-lg font-semibold text-gray-900">{{ job.experience }}</div>
            </div>
            <div class="border rounded-lg p-4">
              <div class="text-sm font-medium text-gray-500">Posted</div>
              <div class="mt-1 text-lg font-semibold text-gray-900">{{ job.postedAt }}</div>
            </div>
            <div class="border rounded-lg p-4">
              <div class="text-sm font-medium text-gray-500">Applications</div>
              <div class="mt-1 text-lg font-semibold text-gray-900">{{ job.applications }}</div>
            </div>
          </div>

          <!-- Job description -->
          <div class="mt-8">
            <h2 class="text-xl font-bold text-gray-900">About the role</h2>
            <div class="mt-4 prose prose-primary max-w-none">
              <p>{{ job.description }}</p>
            </div>
          </div>

          <!-- Requirements -->
          <div class="mt-8">
            <h2 class="text-xl font-bold text-gray-900">Requirements</h2>
            <ul class="mt-4 space-y-2">
              <li v-for="(requirement, index) in job.requirements" :key="index" class="flex items-start">
                <i class="fas fa-check text-green-500 mt-1 mr-2"></i>
                <span>{{ requirement }}</span>
              </li>
            </ul>
          </div>

          <!-- Skills -->
          <div class="mt-8">
            <h2 class="text-xl font-bold text-gray-900">Required Skills</h2>
            <div class="mt-4 flex flex-wrap gap-2">
              <span
                v-for="skill in job.skills"
                :key="skill"
                class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium bg-gray-100 text-gray-800"
              >
                {{ skill }}
              </span>
            </div>
          </div>

          <!-- Apply button -->
          <div class="mt-8 flex justify-end">
            <button
              type="button"
              class="inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
            >
              Apply for this position
            </button>
          </div>
        </div>
      </div>

      <!-- Similar jobs -->
      <div class="mt-12">
        <h2 class="text-2xl font-bold text-gray-900">Similar Jobs</h2>
        <div class="mt-6 grid gap-6 lg:grid-cols-2">
          <div v-for="similarJob in similarJobs" :key="similarJob.id" class="bg-white shadow rounded-lg p-6">
            <div class="flex items-center justify-between">
              <div class="flex items-center">
                <div class="flex-shrink-0 h-12 w-12">
                  <img :src="similarJob.companyLogo" :alt="similarJob.companyName" class="h-12 w-12 rounded-full" />
                </div>
                <div class="ml-4">
                  <h3 class="text-lg font-medium text-gray-900">
                    {{ similarJob.title }}
                  </h3>
                  <div class="text-sm text-gray-500">
                    {{ similarJob.companyName }} • {{ similarJob.location }}
                  </div>
                </div>
              </div>
              <router-link
                :to="{ name: 'JobDetails', params: { id: similarJob.id }}"
                class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-primary-600 bg-primary-50 hover:bg-primary-100"
              >
                View Details
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// Sample job data
const job = {
  id: 1,
  title: 'Senior Frontend Developer',
  companyName: 'TechCorp',
  companyLogo: 'https://via.placeholder.com/150',
  location: 'Remote',
  type: 'Full-time',
  salary: '$120,000 - $160,000',
  experience: '5+ years',
  postedAt: '2 days ago',
  applications: '24 applied',
  description: 'We are looking for a senior frontend developer with strong experience in Vue.js and TypeScript. The ideal candidate will have a passion for building beautiful, responsive, and performant web applications.',
  requirements: [
    'Bachelor\'s degree in Computer Science or related field',
    '5+ years of experience with modern JavaScript frameworks',
    'Strong understanding of TypeScript and static typing',
    'Experience with state management solutions (Vuex/Pinia)',
    'Knowledge of modern frontend build tools and workflows',
    'Excellent problem-solving and communication skills'
  ],
  skills: ['Vue.js', 'TypeScript', 'Tailwind CSS', 'REST APIs', 'Git', 'Jest']
}

// Sample similar jobs
const similarJobs = [
  {
    id: 2,
    title: 'Frontend Developer',
    companyName: 'WebTech Inc',
    companyLogo: 'https://via.placeholder.com/150',
    location: 'San Francisco, CA',
    type: 'Full-time'
  },
  {
    id: 3,
    title: 'UI Engineer',
    companyName: 'DesignCo',
    companyLogo: 'https://via.placeholder.com/150',
    location: 'Remote',
    type: 'Contract'
  }
]
</script>
