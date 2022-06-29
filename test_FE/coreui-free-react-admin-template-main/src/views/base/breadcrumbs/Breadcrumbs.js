import React from 'react'
import { CCol, CRow, CFormLabel, CFormInput, CFormTextarea, CButton } from '@coreui/react'

const Breadcrumbs = () => {
  return (
    <CRow>
      <CCol xs={12}>
        <CRow className="mb-3">
          <CFormLabel htmlFor="inputTitle" className="col-sm-2 col-form-label">
            Title
          </CFormLabel>
          <CCol sm={10}>
            <CFormInput type="text" id="inputTitle" />
          </CCol>
        </CRow>
        <CRow className="mb-3">
          <CFormLabel htmlFor="inputContent">Content</CFormLabel>
          <CFormTextarea id="inpuctContent" rows="3"></CFormTextarea>
        </CRow>
        <CRow className="mb-3">
          <CFormLabel htmlFor="inputCategory" className="col-sm-2 col-form-label">
            Category
          </CFormLabel>
          <CCol sm={10}>
            <CFormInput type="text" id="inputCategory" />
          </CCol>
        </CRow>
        <CButton color="primary">Publish</CButton>&nbsp;<CButton color="secondary">Draft</CButton>
      </CCol>
    </CRow>
  )
}

export default Breadcrumbs
