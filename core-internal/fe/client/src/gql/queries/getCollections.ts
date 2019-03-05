import gql from 'graphql-tag';

const GET_COLLECTIONS = gql`
  query getCollections {
    getCollections {
      id
      name
      host
      fieldValues {
        valueType
        fieldName
        isIndexed
        isUnique
      }
    }
  }
`;
export default GET_COLLECTIONS;
