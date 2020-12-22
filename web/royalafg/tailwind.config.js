module.exports = {
  future : {
    removeDeprecatedGapUtilities : true,
    purgeLayersByDefault : true,
  },
  purge : [],
  theme : {
    fontFamily : {"sans" : [ "Poppins", "sans-serif" ]},
    extend : {},
  },
  variants : {},
  plugins : [
    require('@tailwindcss/custom-forms'),
  ],
}
