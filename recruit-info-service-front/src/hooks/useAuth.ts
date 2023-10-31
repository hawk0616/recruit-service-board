import axios from 'axios'
import { useEffect, useState } from 'react'
import { useRouter } from 'next/router'
import { CsrfToken } from '../types/auth'
import { Credential } from '../types/user'
import { useError } from './useError'

export const useAuth = () => {
  const router = useRouter()
  const { ErrorHandling } = useError()
  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [error, setError] = useState(null)

  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false)

   // conform CSRF from cookies
   useEffect(() => {
    axios.defaults.withCredentials = true
    const getCsrfToken = async () => {
      const { data } = await axios.get<CsrfToken>(
        `${process.env.NEXT_PUBLIC_API_URL}/csrf`,
      )
      axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
    }
    getCsrfToken()
  }, [])

  const login = async (user: Credential) => {
    setIsLoading(true)
    try {
      await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/login`, user)
      setIsAuthenticated(true)
      router.push('/company')
    } catch (err: any) {
      if (err.response.data.message) {
        ErrorHandling(err.response.data.message)
      } else {
        ErrorHandling(err.response.data)
      }
      setError(err)
    } finally {
      setIsLoading(false)
    }
  }

  const signup = async (user: Credential) => {
    setIsLoading(true)
    try {
      await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/signup`, user)
      setIsAuthenticated(true)
      router.push('/company')
    } catch (err: any) {
      if (err.response.data.message) {
        ErrorHandling(err.response.data.message)
      } else {
        ErrorHandling(err.response.data)
      }
      setError(err)
    } finally {
      setIsLoading(false)
    }
  }

  const logout = async () => {
    setIsLoading(true)
    try {
      await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/logout`)
      router.push('/')
    } catch (err: any) {
      if (err.response.data.message) {
        ErrorHandling(err.response.data.message)
      } else {
        ErrorHandling(err.response.data)
      }
      setError(err)
    } finally {
      setIsLoading(false)
    }
  }

  return { login, signup, logout, isLoading, error, isAuthenticated }
}