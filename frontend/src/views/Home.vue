<template>
  <div class="relative bg-white overflow-hidden">
    <!-- Hero section -->
    <div class="relative pb-16 sm:pb-24">
      <main class="mt-16 mx-auto max-w-7xl px-4 sm:mt-24">
        <div class="text-center">
          <h1 class="text-4xl tracking-tight font-extrabold text-gray-900 sm:text-5xl md:text-6xl">
            <span class="block">Professional networking for</span>
            <span class="block text-primary-600">open source developers</span>
          </h1>
          <p class="mt-3 max-w-md mx-auto text-base text-gray-500 sm:text-lg md:mt-5 md:text-xl md:max-w-3xl">
            Connect with other developers, find job opportunities, and contribute to open source projects.
          </p>
          <div class="mt-5 max-w-md mx-auto sm:flex sm:justify-center md:mt-8">
            <div class="rounded-md shadow">
              <router-link
                to="/jobs"
                class="w-full flex items-center justify-center px-8 py-3 border border-transparent text-base font-medium rounded-md text-white bg-primary-600 hover:bg-primary-700 md:py-4 md:text-lg md:px-10"
              >
                Find Jobs
              </router-link>
            </div>
            <div class="mt-3 rounded-md shadow sm:mt-0 sm:ml-3">
              <router-link
                to="/network"
                class="w-full flex items-center justify-center px-8 py-3 border border-transparent text-base font-medium rounded-md text-primary-600 bg-white hover:bg-gray-50 md:py-4 md:text-lg md:px-10"
              >
                Browse Network
              </router-link>
            </div>
          </div>
        </div>
      </main>
    </div>

    <!-- Feature section -->
    <div class="relative bg-gray-50 py-16 sm:py-24 lg:py-32">
      <div class="mx-auto max-w-md px-4 text-center sm:max-w-3xl sm:px-6 lg:max-w-7xl lg:px-8">
        <h2 class="text-base font-semibold uppercase tracking-wider text-primary-600">Everything you need</h2>
        <p class="mt-2 text-3xl font-extrabold text-gray-900 tracking-tight sm:text-4xl">
          Features that empower developers
        </p>
        <p class="mx-auto mt-5 max-w-prose text-xl text-gray-500">
          Everything you need to grow your career and connect with the open source community.
        </p>
        <div class="mt-12">
          <div class="grid grid-cols-1 gap-8 sm:grid-cols-2 lg:grid-cols-3">
            <div class="pt-6">
              <div class="flow-root bg-white rounded-lg px-6 pb-8">
                <div class="-mt-6">
                  <div>
                    <span class="inline-flex items-center justify-center p-3 bg-primary-500 rounded-md shadow-lg">
                      <i class="fas fa-network-wired text-white text-2xl"></i>
                    </span>
                  </div>
                  <h3 class="mt-8 text-lg font-medium text-gray-900 tracking-tight">Decentralized Network</h3>
                  <p class="mt-5 text-base text-gray-500">
                    Connect with developers across different platforms using ActivityPub protocol.
                  </p>
                </div>
              </div>
            </div>

            <div class="pt-6">
              <div class="flow-root bg-white rounded-lg px-6 pb-8">
                <div class="-mt-6">
                  <div>
                    <span class="inline-flex items-center justify-center p-3 bg-primary-500 rounded-md shadow-lg">
                      <i class="fas fa-briefcase text-white text-2xl"></i>
                    </span>
                  </div>
                  <h3 class="mt-8 text-lg font-medium text-gray-900 tracking-tight">Job Opportunities</h3>
                  <p class="mt-5 text-base text-gray-500">
                    Find remote-friendly jobs at companies that value open source contributions.
                  </p>
                </div>
              </div>
            </div>

            <div class="pt-6">
              <div class="flow-root bg-white rounded-lg px-6 pb-8">
                <div class="-mt-6">
                  <div>
                    <span class="inline-flex items-center justify-center p-3 bg-primary-500 rounded-md shadow-lg">
                      <i class="fas fa-code-branch text-white text-2xl"></i>
                    </span>
                  </div>
                  <h3 class="mt-8 text-lg font-medium text-gray-900 tracking-tight">Project Discovery</h3>
                  <p class="mt-5 text-base text-gray-500">
                    Discover open source projects that match your interests and skills.
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- CTA section -->
    <div class="bg-primary-700">
      <div class="max-w-2xl mx-auto text-center py-16 px-4 sm:py-20 sm:px-6 lg:px-8">
        <h2 class="text-3xl font-extrabold text-white sm:text-4xl">
          <span class="block">Ready to dive in?</span>
          <span class="block">Join the network today.</span>
        </h2>
        <p class="mt-4 text-lg leading-6 text-primary-200">
          Start connecting with developers, finding opportunities, and growing your career.
        </p>
        <a
          href="#"
          class="mt-8 w-full inline-flex items-center justify-center px-5 py-3 border border-transparent text-base font-medium rounded-md text-primary-600 bg-white hover:bg-primary-50 sm:w-auto"
        >
          Sign up for free
        </a>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// Component logic here
</script>
