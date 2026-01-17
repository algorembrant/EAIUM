import { render, screen } from '@testing-library/react'
import { describe, it, expect } from 'vitest'
import App from './App'

describe('App', () => {
    it('renders without crashing', () => {
        render(<App />)
        // Adjust this expectation based on actual App content, 
        // for now just checking if it renders *something* or doesn't throw
        // detailed checks depend on what App actually renders (e.g. login screen or dashboard)
        const appContainer = document.querySelector('#root') // Note: this might not work dependent on how render works, better to query by text
    })
})
