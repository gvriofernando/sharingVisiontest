import React, { useState, useEffect } from 'react'
import {
  CRow,
  CPaginationItem,
  CPagination,
  CCard,
  CCardBody,
  CCardTitle,
  CCardSubtitle,
  CCardText,
} from '@coreui/react'
import axios from 'axios'

const Cards = () => {
  const [posts, setPosts] = useState([])
  const [loading, setLoading] = useState(false)
  const [currentPage, setCurrentPage] = useState(1)
  const [postsPerPage] = useState(3)

  useEffect(() => {
    const fetchPosts = async () => {
      setLoading(true)
      const res = await axios.get(
        `http://localhost:8080/articles/${postsPerPage}/${(currentPage - 1) * postsPerPage}`,
      )
      setPosts(res.data.result)
      setLoading(false)
    }

    fetchPosts()
  }, [])

  if (loading && posts.length === 0) {
    return <h2>Loading...</h2>
  }

  const changePages = (event, numPage) => {
    console.log(event)
    setCurrentPage(numPage)

    const fetchPosts = async () => {
      setLoading(true)
      const res = await axios.get(
        `http://localhost:8080/articles/${postsPerPage}/${(currentPage - 1) * postsPerPage}`,
      )
      setPosts(res.data.result)
      setLoading(false)
    }

    fetchPosts()
  }

  return (
    <div>
      <div>
        <CRow>
          <CPagination aria-label="Page navigation example">
            <CPaginationItem onClick={(event) => changePages(event, 1)}>1</CPaginationItem>
            <CPaginationItem onClick={(event) => changePages(event, 2)}>2</CPaginationItem>
            <CPaginationItem onClick={(event) => changePages(event, 3)}>3</CPaginationItem>
          </CPagination>
        </CRow>
      </div>
      <div>
        {posts.map((post) => (
          <CCard style={{ width: '18rem' }} key={post.id}>
            <CCardBody>
              <CCardTitle>{post.title}</CCardTitle>
              <CCardSubtitle className="mb-2 text-medium-emphasis">{post.category}</CCardSubtitle>
              <CCardText>{post.content}</CCardText>
            </CCardBody>
          </CCard>
        ))}
      </div>
    </div>
  )
}

export default Cards
