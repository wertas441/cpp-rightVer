// Copyright © 2023 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React, { useCallback, useContext, useEffect, useRef } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import {
  Route,
  Routes,
  Outlet,
  createBrowserRouter,
  RouterProvider,
  ScrollRestoration,
  useLocation,
} from 'react-router-dom'
import classnames from 'classnames'

import { ToastContainer } from '@ttn-lw/components/toast'
import Breadcrumbs from '@ttn-lw/components/breadcrumbs'
import HeaderComponent from '@ttn-lw/components/header'

import GenericNotFound from '@ttn-lw/lib/components/full-view-error/not-found'
import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'
import ErrorView from '@ttn-lw/lib/components/error-view'
import WithAuth from '@ttn-lw/lib/components/with-auth'
import FullViewError, { FullViewErrorInner } from '@ttn-lw/lib/components/full-view-error'

import Header from '@console/containers/header'
import LogBackInModal from '@console/containers/log-back-in-modal'
import Sidebar from '@console/containers/sidebar'
import EventSplitFrameContext from '@console/containers/event-split-frame/context'
import Logo from '@console/containers/logo'
import { SidebarContextProvider } from '@console/containers/sidebar/context'

import OverviewRoutes from '@console/views/overview'
import Applications from '@console/views/applications'
import Gateways from '@console/views/gateways'

import { setStatusOnline } from '@ttn-lw/lib/store/actions/status'
import { selectStatusStore } from '@ttn-lw/lib/store/selectors/status'
import {
  selectApplicationSiteName,
  selectApplicationSiteTitle,
  selectPageData,
} from '@ttn-lw/lib/selectors/env'

import {
  selectUser,
  selectUserFetching,
  selectUserError,
  selectUserRights,
  selectUserIsAdmin,
} from '@console/store/selectors/user'

import style from './app.styl'

const errorRender = error => (
  <FullViewError error={error} header={<Header alwaysShowLogo Logo={Logo} />} />
)
const getScrollRestorationKey = location => {
  // Preserve scroll position only when necessary.
  // E.g. we don't want to scroll to top when changing tabs of a table,
  // but we do want to scroll to top when changing pages.
  const { pathname, search } = location
  const params = new URLSearchParams(search)
  const page = params.get('page')

  return `${pathname}${page ? `?page=${page}` : ''}`
}

const Layout = () => {
  const { search } = useLocation()
  const page = new URLSearchParams(search).get('page')
  const user = useSelector(selectUser)
  const fetching = useSelector(selectUserFetching)
  const error = useSelector(selectUserError)
  const rights = useSelector(selectUserRights)
  const isAdmin = useSelector(selectUserIsAdmin)
  const siteTitle = selectApplicationSiteTitle()
  const siteName = selectApplicationSiteName()
  const main = useRef()

  const { height: splitFrameHeight } = useContext(EventSplitFrameContext)

  useEffect(() => {
    if (main.current) {
      main.current.scrollTop = 0
    }
  }, [page])

  return (
    <SidebarContextProvider>
      <ScrollRestoration getKey={getScrollRestorationKey} />
      <ErrorView errorRender={errorRender}>
        <div className={style.container}>
          <IntlHelmet
            titleTemplate={`%s - ${siteTitle ? `${siteTitle} - ` : ''}${siteName}`}
            defaultTitle={siteName}
          />
          <div id="modal-container" />
          <div id="dropdown-container" className="pos-absolute-container" />
          <div className="d-flex">
            <Sidebar />
            <div className="w-full h-vh d-flex direction-column">
              <Header />
              <main
                className={classnames(style.main, 'd-flex', 'flex-column', 'h-full', 'flex-grow')}
                ref={main}
              >
                <WithAuth
                  user={user}
                  fetching={fetching}
                  userError={error}
                  errorComponent={FullViewErrorInner}
                  rights={rights}
                  isAdmin={isAdmin}
                >
                  <div className={style.content}>
                    <div
                      className={style.stage}
                      id="stage"
                      style={{ paddingBottom: splitFrameHeight }}
                    >
                      <Outlet />
                    </div>
                  </div>
                </WithAuth>
              </main>
            </div>
          </div>

          <div id="split-frame" className={style.splitFrame} />
        </div>
      </ErrorView>
    </SidebarContextProvider>
  )
}
const ConsoleRoot = () => {
  const status = useSelector(selectStatusStore)
  const pageData = selectPageData()
  const dispatch = useDispatch()

  const handleConnectionStatusChange = useCallback(
    ({ type }) => {
      dispatch(setStatusOnline(type === 'online'))
    },
    [dispatch],
  )

  useEffect(() => {
    window.addEventListener('online', handleConnectionStatusChange)
    window.addEventListener('offline', handleConnectionStatusChange)
    return () => {
      window.removeEventListener('online', handleConnectionStatusChange)
      window.removeEventListener('offline', handleConnectionStatusChange)
    }
  }, [handleConnectionStatusChange])

  if (pageData && pageData.error) {
    return (
      <FullViewError
        error={pageData.error}
        header={
          <HeaderComponent
            safe
            alwaysShowLogo
            Logo={Logo}
            isSidebarMinimized={false}
            expandSidebar={false}
            toggleSidebarMinimized={false}
          />
        }
        safe
      />
    )
  }

  return (
    <React.Fragment>
      {status.isLoginRequired && <LogBackInModal />}
      <ToastContainer />
      <Breadcrumbs />
      <Routes>
        <Route element={<Layout />}>
          <Route path="/*" Component={OverviewRoutes} />
          <Route path="applications/*" Component={Applications} />
          <Route path="gateways/*" Component={Gateways} />
          <Route path="*" Component={GenericNotFound} />
        </Route>
      </Routes>
    </React.Fragment>
  )
}

const router = createBrowserRouter([{ path: '*', Component: ConsoleRoot }], {
  basename: '/console',
})

const ConsoleApp = () => <RouterProvider router={router} />

export default ConsoleApp
