const getRandomElement = <T>(array: readonly T[]) =>
	array[Math.round(Math.random() * (array.length - 1))]

export { getRandomElement }
