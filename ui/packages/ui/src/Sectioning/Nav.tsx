import { Alignment, AnchorButton, Button, Classes, HTMLSelect, IconName, Navbar } from "@blueprintjs/core";
import { useTranslation } from "react-i18next";
import { Link } from "wouter";
import logo from "./logo.png";


export default ({}) => {
  const { t, i18n } = useTranslation();

  const PageNavs = [
    <NavLink key="sim" href="/simulator" icon="calculator" text={t("nav.simulator")} />,
    <NavLink key="view" href="/viewer" icon="chart" text={t("nav.viewer")} />,
    <NavLink key="sample" href="/sample" icon="helper-management" text={"Sample"} />,
    <AnchorButton
        key="db"
        className={Classes.MINIMAL}
        icon="database"
        href="https://db.gcsim.app"
        target="_blank">
      <span className="hidden min-[798px]:block">{t<string>("nav.teams_db")}</span>
    </AnchorButton>,
    <NavLink key="about" href="/about" icon="info-sign" text={t("nav.about")} />,
  ];

  return (
    <Navbar>
      <div className="w-full 2xl:mx-auto 2xl:container">
        <Navbar.Group align={Alignment.LEFT}>
          <Navbar.Heading className="!mr-[10px]">
            <Link href="/" className="flex h-[50px] items-center">
              <img src={logo} className="object-scale-down max-h-[75%] m-auto mr-2" />
              <span className="font-medium font-mono">gcsim</span>
            </Link>
          </Navbar.Heading>
        </Navbar.Group>
        <Navbar.Group align={Alignment.LEFT} className="!hidden min-[550px]:!flex">
          <Navbar.Divider />
          {PageNavs}
        </Navbar.Group>
        <Navbar.Group align={Alignment.RIGHT}>
          {/* <Link href="/account">
            <Button minimal={true} icon="user" text="Guest" />
          </Link> */}
          <HTMLSelect
              className="ml-2"
              value={i18n.resolvedLanguage}
              onChange={(e) => i18n.changeLanguage(e.target.value)}>
            <option value="en">{t<string>("nav.english")}</option>
            <option value="zh">{t<string>("nav.chinese")}</option>
            <option value="de">{t<string>("nav.german")}</option>
            <option value="ja">{t<string>("nav.japanese")}</option>
            <option value="es">{t<string>("nav.spanish")}</option>
            <option value="ru">{t<string>("nav.russian")}</option>
          </HTMLSelect>
        </Navbar.Group>
      </div>
    </Navbar>
  );
};

const NavLink = ({ href, icon, text }: { href: string, icon: IconName , text: string }) => {
  return (
    <Link href={href}>
      <Button minimal={true} icon={icon}>
        <span className="hidden min-[798px]:block">{text}</span>
      </Button>
    </Link>
  );
};