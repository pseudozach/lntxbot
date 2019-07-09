/** @format */

import React from 'react'

import CardBack from './back'

// example cards
//  { rank: "A", type: "H" },
// { rank: "10", type: "D" },

const club =
  'M140.693 193.819C81.88 165.724-4.634 199.236 15.835 274.83c19.947 73.668 106.09 57.501 135.879 12.658-12.428 85.025-34.253 98.263-47.334 116.72h138.884c-14.072-19.739-39.624-31.695-54.609-117.885 30.017 44.768 120.253 60.287 139.912-12.317 19.339-71.42-67.833-112.094-125.8-79.762 55.585-37.326 75.547-159.914-29.732-159.914-106.317 0-88.579 124.985-32.342 159.488z'
const heart =
  'M175.04 405.115c-.344-.619-1.819-5.088-3.28-9.934-9.253-30.705-26.432-61.799-55.2-99.912-10.807-14.319-20.98-26.998-48.489-60.437-31.851-38.717-40.975-51.12-50.188-68.228-5.44-10.102-11.078-25.474-12.933-35.263-1.855-9.783-1.885-25.766-.065-34.089 7.403-33.854 37.475-59.395 73.433-62.369 41.193-3.407 76.233 17.403 94.643 56.207l4.112 8.665 3.094-6.845c4.597-10.172 9.025-17.051 16.499-25.635 19.028-21.853 41.673-32.558 68.954-32.596 12.22-.016 19.222.998 29.636 4.293 15.573 4.928 27.297 12.605 37.738 24.713 27.21 31.553 23.988 74.54-9.532 127.188-7.82 12.28-21.928 30.452-41.33 53.23-22.343 26.233-33.26 39.564-44.438 54.266-26.791 35.24-44.81 67.26-54.72 97.243-1.512 4.574-3.096 8.836-3.52 9.472-.985 1.475-3.602 1.493-4.414.03z'
const spade =
  'M67.299 130.308c-122.08 124.722-2.408 215.191 94.114 155.646-18.225 64.968-28.248 70.08-39.312 87.19h104.214c-13.377-17.277-27.426-22.222-41.66-86.82 98.547 56.61 202.41-42.114 93.06-156.016C206.202 67.753 178.438 10.643 172.508 2.534 166.818 9.291 133.3 68.914 67.3 130.308z'
const diamond =
  'M175.04 29.435C136.342 93.832 92.777 155.78 48.518 215.275c46.345 59.496 90.604 121.442 126.522 185.84 37.307-65.623 78.091-128.182 126.522-185.84-47.04-59.496-89.91-122.668-126.522-185.84z'

const suitmap = {
  C: club,
  H: heart,
  S: spade,
  D: diamond
}

const CARD_WIDTH = 1800
const CARD_HEIGHT = CARD_WIDTH * 1.453

export {CardBack}

export default ({rank, type}) => {
  return (
    <svg className="card" viewBox={`0 0 ${CARD_WIDTH} ${CARD_HEIGHT}`}>
      <g>
        <rect
          x="0"
          y="0"
          rx="100"
          ry="100"
          width={CARD_WIDTH}
          height={CARD_HEIGHT}
        />
      </g>
      <g className={`${type} x${rank}`}>
        <text>{rank}</text>
      </g>
      <g className={`${type} x${rank}`}>
        <path d={suitmap[type]} />
      </g>
      <g className={`${type}`}>
        <path d={suitmap[type]} />
      </g>
    </svg>
  )
}
