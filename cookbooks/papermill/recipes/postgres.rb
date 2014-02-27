apt_repository 'pgdg' do
  uri          'http://apt.postgresql.org/pub/repos/apt/'
  distribution 'precise-pgdg'
  components   ['main']
  key          'https://www.postgresql.org/media/keys/ACCC4CF8.asc'
end

package 'postgresql-9.3'
package 'postgresql-contrib-9.3'
