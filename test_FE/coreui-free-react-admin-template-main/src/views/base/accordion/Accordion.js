/* eslint-disable */
import React, { useState, useEffect } from 'react'
import { 
  CNav, 
  CNavItem, 
  CNavLink,
  CTable,
  CTableHead,
  CTableBody,
  CTableRow,
  CTableHeaderCell,
  CTableDataCell,
} from '@coreui/react'
import CIcon from '@coreui/icons-react'
import {
  cilPencil,
  cilTrash,
} from '@coreui/icons'
import axios from 'axios'

const Accordion = () => {
  const [posts, setPosts] = useState([])
  const [loading, setLoading] = useState(false)
  const [currentPage, setCurrentPage] = useState("publish")

  useEffect(() => {
    const fetchPosts = async () => {
      setLoading(true)
      const res = await axios.get(
        `http://localhost:8080/articleKhusus/${currentPage}`,
      )
      setPosts(res.data.result)
      setLoading(false)
    }

    fetchPosts()
  }, [])

  if (loading && posts.length === 0) {
    return <h2>Loading...</h2>
  }

  const changePages = (event, status) => {
    setCurrentPage(status)
  
    const fetchPosts = async () => {
      setLoading(true)
      const res = await axios.get(
        `http://localhost:8080/articleKhusus/${currentPage}`,
      )
      setPosts(res.data.result)
      setLoading(false)
    }
  
    fetchPosts()
  }

  return (
    <div className="pageAllPosts">
      <div className="navBar">
        <CNav variant="tabs">
          <CNavItem>
            <CNavLink onClick={(event) => changePages(event, "publish")} className="">Published</CNavLink>
          </CNavItem>
          <CNavItem>
            <CNavLink onClick={(event) => changePages(event, "draft")} className="">Drafts</CNavLink>
          </CNavItem>
          <CNavItem>
            <CNavLink onClick={(event) => changePages(event, "thrased")} className="">Thrased</CNavLink>
          </CNavItem>
        </CNav>
      </div>
      <div>
        <CTable>
          <CTableHead>
            <CTableRow>
              <CTableHeaderCell scope="col">No</CTableHeaderCell>
              <CTableHeaderCell scope="col">Title</CTableHeaderCell>
              <CTableHeaderCell scope="col">Category</CTableHeaderCell>
              <CTableHeaderCell scope="col">Action</CTableHeaderCell>
            </CTableRow>
          </CTableHead>
          <CTableBody>
            {posts.map((post) => (
              <CTableRow key={post.id}>
                <CTableHeaderCell scope="row">{post.id}</CTableHeaderCell>
                <CTableDataCell>{post.title}</CTableDataCell>
                <CTableDataCell>{post.category}</CTableDataCell>
                <CTableDataCell>
                  <CIcon icon={cilPencil} size="lg"/>
                  &nbsp;&nbsp;&nbsp;
                  <CIcon icon={cilTrash} size="lg"/>
                </CTableDataCell>
              </CTableRow>
            ))}
          </CTableBody>
        </CTable>
      </div>
    </div>
  )
}

export default Accordion
