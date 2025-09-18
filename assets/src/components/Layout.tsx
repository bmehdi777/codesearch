import { Outlet } from "react-router";

function Layout() {
  return (
    <div className="flex w-dvw h-dvh">
      <Outlet />
    </div>
  );
}

export default Layout;
