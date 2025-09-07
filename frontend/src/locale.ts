import { configureLocalization } from "@lit/localize";
import { sourceLocale, targetLocales } from "./generated/locale-codes.js";

export function initLocalize() {
  const { getLocale, setLocale } = configureLocalization({
    sourceLocale,
    targetLocales,
    loadLocale: (locale) =>
      ({
        de: () => import("./generated/locales/de"),
        fa: () => import("./generated/locales/fa"),
        fi: () => import("./generated/locales/fi"),
        fr: () => import("./generated/locales/fr"),
        it: () => import("./generated/locales/it"),
        "nb-NO": () => import("./generated/locales/nb-NO"),
        nl: () => import("./generated/locales/nl"),
        pl: () => import("./generated/locales/pl"),
        "pt-BR": () => import("./generated/locales/pt-BR"),
        ru: () => import("./generated/locales/ru"),
        tr: () => import("./generated/locales/tr"),
        "zh-Hans": () => import("./generated/locales/zh-Hans"),
      })[locale](),
  });

  setLocale(document.documentElement.lang);
  globalThis.setLocale = setLocale;
  globalThis.getLocale = getLocale;
}
